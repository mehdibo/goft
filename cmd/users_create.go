package cmd

import (
	"errors"
	"fmt"
	"goft/pkg/ftapi"
	"strconv"

	"github.com/spf13/cobra"
)

// NewUserCreateCmd create the users create cmd
func NewUserCreateCmd(api *ftapi.APIInterface) *cobra.Command {
	cmd := cobra.Command{
		Use:   "create email first_name last_name kind campus_id",
		Short: "Create a new user",
		Long: `This command requires the Advanced tutor role
No password is set, the user should reset his password using the web interface.

kind must be either admin, student or external.`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 5 {
				return fmt.Errorf("accepts 5 arg(s), received %d", len(args))
			}
			if args[3] != "admin" && args[3] != "student" && args[3] != "external" {
				return errors.New("kind must be admin, student or external")
			}
			campusID, err := strconv.Atoi(args[4])
			if err != nil || campusID <= 0 {
				return errors.New("invalid campus_id")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			campusID, _ := strconv.Atoi(args[4])
			user := ftapi.User{
				Email:     args[0],
				FirstName: args[1],
				LastName:  args[2],
				Kind:      args[3],
				CampusID:  campusID,
			}
			login, _ := cmd.Flags().GetString("login")
			if login != "" {
				user.Login = login
			}
			err := (*api).CreateUser(&user)
			// TODO: stop printing errors and let cobra handle the return
			if err != nil {
				return err
			}
			_, _ = fmt.Fprint(cmd.OutOrStdout(), "User created\n")
			return nil
		},
	}
	cmd.Flags().String("login", "", "Set a custom login, leave empty to let the intra generate one")
	return &cmd
}
var userCreateCmd = NewUserCreateCmd(&API)

func init() {
	usersCmd.AddCommand(userCreateCmd)
}
