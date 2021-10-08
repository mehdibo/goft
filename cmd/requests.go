package cmd

import (
	"github.com/spf13/cobra"
)

// NewRequestsCmd creates requests cmd
func NewRequestsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "requests",
		Short: "Send raw requests to the Intranet's API",
	}
}
var requestsCmd = NewRequestsCmd()

func init() {
	rootCmd.AddCommand(requestsCmd)
}
