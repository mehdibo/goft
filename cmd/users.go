package cmd

import (
	"github.com/spf13/cobra"
)

// usersCmd represents the users command
func NewUsersCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "users",
		Short: "Interact with the user entity",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

}
var usersCmd = NewUsersCmd()
func init() {
	rootCmd.AddCommand(usersCmd)
}
