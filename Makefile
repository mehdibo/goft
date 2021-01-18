GOCMD=go
NAME=goft

all: $(NAME)

$(NAME): vendor cmd/*.go main.go
	@echo "Building binary..."
	$(GOCMD) build -o $(NAME) -v

vendor: go.mod
	@echo "Installing dependencies..."
	$(GOCMD) mod vendor

.PHONY: fclean
fclean: clean
	-rm -rf vendor

.PHONY: clean
clean:
	-rm -f $(NAME)
	-rm -f $(NAME)-darwin-*
	-rm -f $(NAME)-linux-*

.PHONY: re
re: fclean all

.PHONY: test
test:
	go test -v ./...

.PHONY: lint
lint: vendor
	golint -set_exit_status ./cmd/... ./pkg/...

.PHONY: all_platforms
all_platforms: vendor
	for GOOS in darwin linux; do \
        for GOARCH in 386 amd64; do \
			export GOOS $GOOS ; \
			export GOARCH $GOARCH; \
			echo "Building binary for $$GOOS - $$GOARCH" ; \
			go build -o $(NAME)-$$GOOS-$$GOARCH; \
		done \
    done