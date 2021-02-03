package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goft/pkg/ftapi"
)

// NewListUsersCmd create new list users cmd
func NewListUsersCmd(api *ftapi.APIInterface) *cobra.Command {
	cmd := cobra.Command{
		Use:   "ls",
		Short: "List users",
		RunE: func(cmd *cobra.Command, args []string) error {
			filters := map[string]string{
				"primary_campus_id": "16",
			}
			sort := map[string]string{
				"id": "desc",
			}
			users, err := (*api).GetUsers(1, &filters, &sort)
			if err != nil {
				return err
			}
			fmt.Println(users)
			return nil
		},
	}
	return &cmd
}
var listUsersCmd = NewListUsersCmd(&API)

func init() {
	usersCmd.AddCommand(listUsersCmd)
}
