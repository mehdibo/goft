package cmd

import (
	"goft/pkg/ftapi"

	"github.com/spf13/cobra"
)

// NewLogoutCmd create the get user cmd
func NewLogoutCmd(api *ftapi.APIInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "logout",
		Short: "Log out of a 42intra host",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Print("[Executed] goft auth logout")
			return nil
		},
	}
}

var logoutCmd = NewLogoutCmd(&API)

func init() {
	authCmd.AddCommand(logoutCmd)
}
