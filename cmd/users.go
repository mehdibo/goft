package cmd

import (
	"github.com/spf13/cobra"
)

// NewUsersCmd Creates new users command
func NewUsersCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "users",
		Short: "Interact with the user entity",
	}

}
var usersCmd = NewUsersCmd()
func init() {
	rootCmd.AddCommand(usersCmd)
}
