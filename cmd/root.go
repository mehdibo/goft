package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goft/pkg/ftapi"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
// API is used to interact with the 42 API
var API ftapi.APIInterface
// NewRootCmd Create new root command
func NewRootCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "goft",
		Short: "CLI tool to interact with 42's API",
	}
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goft.yaml)")
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
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".goft" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".goft")
	}

	viper.SetDefault("token_endpoint", "https://api.intra.42.fr/oauth/token")
	viper.SetDefault("api_endpoint", "https://api.intra.42.fr/v2")
	viper.SetDefault("scopes", []string{"profile"})

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	requiredConfigs := []string{
		"client_id",
		"client_secret",
	}
	for _, requiredConfig := range requiredConfigs {
		if viper.GetString(requiredConfig) == "" {
			_, _ = fmt.Fprintf(rootCmd.OutOrStderr(), "%s is required but not set in the config file\n", requiredConfig)
			// This will interfere with unit tests
			// TODO: look for a workaround
			os.Exit(1)
		}
	}

	API = ftapi.NewFromCredentials(viper.GetString("api_endpoint"), &clientcredentials.Config{
		ClientID:       viper.GetString("client_id"),
		ClientSecret:   viper.GetString("client_secret"),
		TokenURL:       viper.GetString("token_endpoint"),
		Scopes:         viper.GetStringSlice("scopes"),
		EndpointParams: nil,
		AuthStyle:      oauth2.AuthStyleInParams,
	})
}
