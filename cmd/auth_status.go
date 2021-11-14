package cmd

import (
	"goft/pkg/ftapi"

	"github.com/spf13/cobra"
)

// NewStatusCmd create the get user cmd
func NewStatusCmd(api *ftapi.APIInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Status stored authentication credentials",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Print("[Executed] goft auth status")
			return nil
		},
	}
}

var statusCmd = NewStatusCmd(&API)

func init() {
	authCmd.AddCommand(statusCmd)
}
