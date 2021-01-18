package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// NewRootCmd Create new root command
func NewRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "go_ft",
		Short: "CLI tool to interact with 42's API",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
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

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go_ft.yaml)")
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

		// Search config in home directory with name ".go_ft" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".go_ft")
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
}
