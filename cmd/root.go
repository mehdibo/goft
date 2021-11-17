package cmd

import (
	"context"
	"fmt"
	"goft/pkg/ftapi"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/cobra"
	"golang.org/x/oauth2"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	// API is used to interact with the 42 API
	API ftapi.APIInterface
	// Version the current used version
	Version = "development-build"
	token   *oauth2.Token
)

// NewRootCmd Create new root command
func NewRootCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "goft",
		Short: "CLI tool to interact with 42's API",
	}
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/goft/secret.yml)")
	cmd.Version = Version
	return &cmd
}

var rootCmd = NewRootCmd()

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		dir, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		dir = filepath.Join(dir, ".config", "goft")
		if err := os.MkdirAll(dir, 700); err != nil {
			os.Exit(1)
		}

		// Search config in home directory with name ".goft" (without extension).
		viper.AddConfigPath(dir)
		viper.SetConfigName("secret")
	}

	// Load config from env variables if available
	viper.SetConfigType("yml")
	viper.AutomaticEnv()

	viper.SetDefault("auth_endpoint", "https://api.intra.42.fr/oauth/authorize")
	viper.SetDefault("token_endpoint", "https://api.intra.42.fr/oauth/token")
	viper.SetDefault("api_endpoint", "https://api.intra.42.fr/v2")
	viper.SetDefault("scopes", []string{"profile"})

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if cfgFile != "" {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}
	}

	requiredConfigs := []string{
		"client_id",
		"client_secret",
	}
	for _, requiredConfig := range requiredConfigs {
		if viper.GetString(requiredConfig) == "" {
			_, _ = fmt.Fprintf(rootCmd.OutOrStderr(), "%s is required but not set in the config file\n", requiredConfig)
			os.Exit(1)
		}
	}
	confIntra := &oauth2.Config{
		ClientID:     viper.GetString("client_id"),
		ClientSecret: viper.GetString("client_secret"),
		Scopes:       viper.GetStringSlice("scopes"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  viper.GetString("auth_endpoint"),
			TokenURL: viper.GetString("token_endpoint"),
		},
		RedirectURL: viper.GetString("redirect_uri"),
	}

	if viper.GetString("access_token") == "" {
		httpServerExitDone := &sync.WaitGroup{}
		httpServerExitDone.Add(1)
		srv := startHttpServer(httpServerExitDone)

		url := confIntra.AuthCodeURL("")

		fmt.Println("Open this URL")
		fmt.Println(url)

		for {
			if token != nil {
				if err := srv.Shutdown(context.Background()); err != nil {
					panic(err)
				}
				httpServerExitDone.Wait()
				break
			}
		}
		viper.Set("access_token", token.AccessToken)
		//		viper.WriteConfig()
	} else {
		token = &oauth2.Token{
			AccessToken: viper.GetString("access_token"),
		}
	}
	ctx := context.Background()
	client := confIntra.Client(ctx, token)
	API = ftapi.New(viper.GetString("api_endpoint"), client)
}

func startHttpServer(wg *sync.WaitGroup) *http.Server {
	srv := &http.Server{Addr: ":4200"}
	http.HandleFunc("/", IntraLoginRHandler)

	go func() {
		defer wg.Done()
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	return srv
}

func IntraLoginRHandler(w http.ResponseWriter, r *http.Request) {
	confIntra := &oauth2.Config{
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
	tok, err := confIntra.Exchange(ctx, code[0])
	if err != nil {
		fmt.Fprintf(w, "OAuth Error:%v", err)
	}
	token = tok
}
