package cmd

import (
	"github.com/spf13/cobra"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth <command>",
	Short: "Login, logout, and refresh your authentication",
	Long:  `Manage goft's authentication state.`,
}

func init() {
	rootCmd.AddCommand(authCmd)
}
