package cmd

import (
	"errors"
	"goft/pkg/ftapi"

	"github.com/spf13/cobra"
)

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
			primaryCampus := user.GetPrimaryCampus()
			cmd.Printf(`Id: %d
Login: %s
Email: %s
First name: %s
Last name: %s
Phone: %s
Image: %s
Is staff: %t
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
				user.PoolMonth,
				user.PoolYear,
			)
			if primaryCampus != nil {
				cmd.Printf("Primary campus: %s\n", primaryCampus.Name)
			}
			return nil
		},
	}
}
var getCmd = NewGetUserCmd(&API)

func init() {
	usersCmd.AddCommand(getCmd)
}
