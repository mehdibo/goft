package cmd

import (
	"bytes"
	"encoding/json"
	"goft/pkg/ftapi"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewLogoutCmd create the get user cmd
func NewLogoutCmd(api *ftapi.APIInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "logout",
		Short: "Log out of a 42intra host",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			configMap := viper.AllSettings()
			delete(configMap, "access_token")
			delete(configMap, "refresh_token")
			delete(configMap, "expiry_time")
			delete(configMap, "token_type")
			encodedConfig, _ := json.MarshalIndent(configMap, "", " ")
			err := viper.ReadConfig(bytes.NewReader(encodedConfig))
			if err != nil {
				return err
			}
			viper.WriteConfig()
			color.Set(color.FgGreen)
			cmd.Println("Logged out as ", viper.GetString("login"))
			color.Set(color.Reset)
			return nil
		},
	}
}

var logoutCmd = NewLogoutCmd(&API)

func init() {
	authCmd.AddCommand(logoutCmd)
}
