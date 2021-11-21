package cmd

import (
	"context"
	"errors"
	"fmt"
	"goft/pkg/ftapi"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"

	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

func genToken(config *oauth2.Config) (*oauth2.Token, error) {
	var token *oauth2.Token

	if viper.GetString("access_token") == "" {
		tokenChan = make(chan *oauth2.Token, 1)

		httpServerExitDone := &sync.WaitGroup{}
		httpServerExitDone.Add(1)
		srv := startHttpServer(httpServerExitDone)

		url := config.AuthCodeURL("")
		browser := "xdg-open"
		if runtime.GOOS == "darwin" {
			browser = "open"
		}
		args := []string{url}
		browser, err := exec.LookPath(browser)
		if err == nil {
			cmd := exec.Command(browser, args...)
			cmd.Stderr = os.Stderr
			err = cmd.Start()
			if err != nil {
				return nil, fmt.Errorf("cannot start command: %v", err)
			}
		} else {
			fmt.Println("Open this URL")
			fmt.Println(url)
		}

		token = <-tokenChan
		if err := srv.Shutdown(context.Background()); err != nil {
			panic(err)
		}
		httpServerExitDone.Wait()

		if login, err := getLogin(config, token); err == nil {
			viper.Set("login", login)
		}

	} else {
		token = &oauth2.Token{
			AccessToken:  viper.GetString("access_token"),
			TokenType:    viper.GetString("token_type"),
			RefreshToken: viper.GetString("refresh_token"),
			Expiry:       viper.GetTime("expiry_time"),
		}
	}

	if token == nil {
		return nil, errors.New("generate token fail")
	}
	return token, nil
}

func getLogin(config *oauth2.Config, token *oauth2.Token) (string, error) {
	ctx := context.Background()
	client := config.Client(ctx, token)
	api := ftapi.New(viper.GetString("api_endpoint"), client)
	user, err := api.GetMe()
	if err != nil {
		return "", err
	}
	return user.Login, nil
}

func saveConfig() {
	token, err := API.GetToken()
	if err != nil {
		log.Println("save token: ", err)
		return
	}
	viper.Set("access_token", token.AccessToken)
	viper.Set("token_type", token.TokenType)
	viper.Set("refresh_token", token.RefreshToken)
	viper.Set("expiry_time", token.Expiry)
	viper.WriteConfig()
}

func startHttpServer(wg *sync.WaitGroup) *http.Server {
	uri := viper.GetString("redirect_uri")
	uri = strings.TrimPrefix(uri, "http://")
	srv := &http.Server{Addr: uri}
	http.HandleFunc("/", LoginRedirectHandler)

	go func() {
		defer wg.Done()
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	return srv
}

func LoginRedirectHandler(w http.ResponseWriter, r *http.Request) {
	config := &oauth2.Config{
		ClientID:     viper.GetString("client_id"),
		ClientSecret: viper.GetString("client_secret"),
		Scopes:       viper.GetStringSlice("scopes"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  viper.GetString("auth_endpoint"),
			TokenURL: viper.GetString("token_endpoint"),
		},
		RedirectURL: viper.GetString("redirect_uri"),
	}

	code := r.URL.Query()["code"]
	if code == nil || len(code) == 0 {
		fmt.Fprint(w, "Invalid Parameter")
		return
	}
	ctx := context.Background()
	tok, err := config.Exchange(ctx, code[0])
	if err != nil {
		fmt.Fprintf(w, "OAuth Error:%v", err)
		return
	}
	tokenChan <- tok
	fmt.Fprintln(w, "Successfully logged in. Please close the tab.")
}
