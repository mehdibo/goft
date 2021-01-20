package cmd

import (
	"github.com/spf13/cobra"
)

// NewClosesCmd creates closes cmd
func NewClosesCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "closes",
		Short: "Manage users' closes",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}
var closesCmd = NewClosesCmd()

func init() {
	rootCmd.AddCommand(closesCmd)
}
