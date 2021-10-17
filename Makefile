GOCMD=go
NAME=goft
INSTALL_PATH=/usr/local/bin
CONFIG_FILE_EXAMPLE=config.example.yml
CONFIG_FILE_TARGET=$$HOME/.goft.yml
VERSION ?= "development-version"
all: $(NAME)

$(NAME): vendor cmd/*.go main.go
	@echo "Building binary..."
	$(GOCMD) build -ldflags "-X 'goft/cmd.Version=$(VERSION)'" -o $(NAME) -v

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
	GOFT_CLIENT_ID=test GOFT_CLIENT_SECRET=test go test -v ./cmd/... ./pkg/...

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
			$(GOCMD) build -ldflags "-X 'goft/cmd.Version=$(VERSION)'" -o $(NAME)-$$GOOS-$$GOARCH; \
		done \
    done

.PHONY: install
install: $(NAME)
	sudo mv $(NAME) $(INSTALL_PATH)
	cp -i $(CONFIG_FILE_EXAMPLE) $(CONFIG_FILE_TARGET)
	@echo "You need to edit $(CONFIG_FILE_TARGET) with your credentials"

.PHONY: uninstall
uninstall:
	sudo rm $(INSTALL_PATH)/$(NAME)
	rm $(CONFIG_FILE_TARGET)