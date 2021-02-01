package cmd

import (
	"errors"
	"fmt"
	"goft/pkg/ftapi"

	"github.com/spf13/cobra"
)

func formatText(user *ftapi.User) string {
	output := fmt.Sprintf(`Id: %d
Login: %s
Email: %s
First name: %s
Last name: %s
Phone: %s
Image: %s
Is staff: %t
Correction points: %d
Pool Month/Year: %s/%s
`,
		user.ID,
		user.Login,
		user.Email,
		user.FirstName,
		user.LastName,
		user.Phone,
		user.ImageURL,
		user.IsStaff,
		user.CorrectionPoints,
		user.PoolMonth,
		user.PoolYear,
	)
	primaryCampus := user.GetPrimaryCampus()
	if primaryCampus != nil {
		output += fmt.Sprintf("Primary campus: %s\n", primaryCampus.Name)
	}
	return output
}

// NewGetUserCmd create the get user cmd
func NewGetUserCmd(api *ftapi.APIInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "get login",
		Short: "Get details about a user",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error{
			user, err := (*api).GetUserByLogin(args[0])
			if err != nil {
				return err
			}
			if user == nil {
				return errors.New("failed getting user")
			}
			cmd.Print(formatText(user))
			return nil
		},
	}
}
var getCmd = NewGetUserCmd(&API)

func init() {
	usersCmd.AddCommand(getCmd)
}
