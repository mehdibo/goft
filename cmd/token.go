package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
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
		fmt.Println("Open this URL")
		fmt.Println(url)

		token = <-tokenChan
		if err := srv.Shutdown(context.Background()); err != nil {
			panic(err)
		}
		httpServerExitDone.Wait()
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
	}
	ctx := context.Background()
	tok, err := config.Exchange(ctx, code[0])
	if err != nil {
		fmt.Fprintf(w, "OAuth Error:%v", err)
	}
	tokenChan <- tok
	fmt.Fprintln(w, "Successfully logged in. Please close the tab.")
}
