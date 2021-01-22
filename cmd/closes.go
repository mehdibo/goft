package cmd

import (
	"github.com/spf13/cobra"
)

// NewClosesCmd creates closes cmd
func NewClosesCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "closes",
		Short: "Manage users' closes",
	}
}
var closesCmd = NewClosesCmd()

func init() {
	rootCmd.AddCommand(closesCmd)
}
