package cmd

import (
	"goft/pkg/ftapi"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// NewLoginCmd create the get user cmd
func NewLoginCmd(api *ftapi.APIInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Authenticate with a 42intra host",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			defer saveConfig()
			user, err := (*api).GetMe()
			if err != nil {
				return err
			}
			color.Set(color.FgGreen)
			cmd.Printf("Logged in as %s\n", user.Login)
			color.Set(color.Reset)
			return nil
		},
	}
}

var loginCmd = NewLoginCmd(&API)

func init() {
	authCmd.AddCommand(loginCmd)
}
