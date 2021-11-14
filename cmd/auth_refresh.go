package cmd

import (
	"goft/pkg/ftapi"

	"github.com/spf13/cobra"
)

// NewRefreshCmd create the get user cmd
func NewRefreshCmd(api *ftapi.APIInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "refresh",
		Short: "Refresh stored authentication credentials",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Print("[Executed] goft auth refresh")
			return nil
		},
	}
}

var refreshCmd = NewRefreshCmd(&API)

func init() {
	authCmd.AddCommand(refreshCmd)
}
