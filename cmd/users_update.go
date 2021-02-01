package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"goft/pkg/ftapi"
	"strings"
)

// NewUpdateUserCmd create the update user cmd
func NewUpdateUserCmd(api *ftapi.APIInterface) *cobra.Command {
	cmd := cobra.Command{
		Use:   "update login",
		Short: "Update a user's data",
		Long: "This command requires the Advanced tutor role",
		Args: func(cmd *cobra.Command, args []string) error {
			exactArgs := cobra.ExactArgs(1)
			if err := exactArgs(cmd, args); err != nil {
				return err
			}
			kind, _ := cmd.LocalFlags().GetString("kind")
			if kind != "" {
				if kind != "admin" && kind != "student" && kind != "external" {
					return errors.New("kind must be admin, student or external")
				}
			}
			// TODO: make sure at least one flag is passed
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			password := ""
			promptPasswd, _ := cmd.LocalFlags().GetBool("password")
			if promptPasswd {
				// TODO: hide password from being typed in the terminal
				reader := bufio.NewReader(cmd.InOrStdin())
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Enter new password: ")
				password, _ = reader.ReadString('\n')
				password = strings.Replace(password, "\n", "", 1)
			}
			email, _ := cmd.LocalFlags().GetString("email")
			fName, _ := cmd.LocalFlags().GetString("first-name")
			lName, _ := cmd.LocalFlags().GetString("last-name")
			kind, _ := cmd.LocalFlags().GetString("kind")


			user := ftapi.User{
				Email:          email,
				FirstName:      fName,
				LastName:       lName,
				Kind:           kind,
				Password:       password,
			}
			err := (*api).UpdateUser(args[0], &user)
			if err != nil {
				return err
			}
			_, _ = fmt.Fprint(cmd.OutOrStdout(), "User updated\n")
			return nil
		},
	}
	cmd.Flags().String("email", "", "New email")
	cmd.Flags().String("first-name", "", "New first name")
	cmd.Flags().String("last-name", "", "New last name")
	cmd.Flags().String("kind", "", "New kind, must be either admin, student or external")
	cmd.Flags().Bool("password", false, "Passing this flag will show a prompt to type in password")
	return &cmd
}
// updateCmd represents the update command
var updateCmd = NewUpdateUserCmd(&API)

func init() {
	usersCmd.AddCommand(updateCmd)
}
