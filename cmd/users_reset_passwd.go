package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goft/pkg/ftapi"
	"github.com/sethvargo/go-password/password"
)

// NewUpdateUserCmd create the update user cmd
func NewResetPasswdCmd(api *ftapi.APIInterface, p password.PasswordGenerator) *cobra.Command {
	cmd := cobra.Command{
		Use:   "reset-passwd login",
		Short: "Send a reset password email to the user",
		Long: "This command requires the Advanced tutor role",
		Args: func(cmd *cobra.Command, args []string) error {
			if err := cobra.ExactArgs(1)(cmd, args); err != nil {
				return err
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			// Get User email
			user, err := (*api).GetUserByLogin(args[0])
			if err != nil {
				return err
			}
			// Generate a new password
			newPass, err := p.Generate(12, 1, 1, false, false)
			if err != nil {
				return err
			}
			// Update the user's password
			newUser := ftapi.User{
				Password: newPass,
			}
			fmt.Println(newPass)
			err = (*api).UpdateUser(user.Login, &newUser)
			if err != nil {
				return err
			}
			// Send the password via email

			return nil
		},
	}
	return &cmd
}

var passwdGenerator, _ = password.NewGenerator(nil)
var resetPasswdCmd = NewResetPasswdCmd(&API, passwdGenerator)

func init() {
	usersCmd.AddCommand(resetPasswdCmd)
}
