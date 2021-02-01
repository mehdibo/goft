package cmd

import (
	"fmt"
	"goft/pkg/ftapi"

	"github.com/spf13/cobra"
)

// listUsersCmd represents the ls command
func NewListUsersCmd(api *ftapi.APIInterface) *cobra.Command {
	cmd := cobra.Command{
		Use:   "ls",
		Short: "List users",
		Run: func(cmd *cobra.Command, args []string) {
			filters := map[string]string{
				"primary_campus_id": "16",
			}
			sort := map[string]string{
				"id": "desc",
			}
			users, err := (*api).GetUsers(1, &filters, &sort)
			fmt.Println(err)
			fmt.Println(users)
		},
	}
	return &cmd
}
var listUsersCmd = NewListUsersCmd(&API)

func init() {
	usersCmd.AddCommand(listUsersCmd)
}
