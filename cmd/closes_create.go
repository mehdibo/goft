package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goft/pkg/ftapi"
)

// NewCloseCreateCmd Create the close create cmd
func NewCloseCreateCmd(api *ftapi.APIInterface) *cobra.Command {
	// login - kind - reason -
	return &cobra.Command{
		Use:   "create login kind reason",
		Short: "Create a new close for a user",
		Long: `This command requires the Basic staff role

kind must be one of the following options: agu, black_hole, deserter, non_admitted, serious_misconduct, social_security or other`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("create called")
			(*api).CreateClose(&ftapi.Close{
				Kind:              "other",
				Reason:            "This is the reason",
				CommunityServices: nil,
				User:              &ftapi.User{
					Login:     "spoody",
				},
			})
		},
	}
}
var createCmd = NewCloseCreateCmd(&API)

func init() {
	closesCmd.AddCommand(createCmd)
}
