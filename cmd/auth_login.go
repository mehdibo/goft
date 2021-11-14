package cmd

import (
	"goft/pkg/ftapi"

	"github.com/spf13/cobra"
)

// NewLoginCmd create the get user cmd
func NewLoginCmd(api *ftapi.APIInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Authenticate with a 42intra host",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Print("[Executed] goft auth login")
			return nil
		},
	}
}

var loginCmd = NewLoginCmd(&API)

func init() {
	authCmd.AddCommand(loginCmd)
}
