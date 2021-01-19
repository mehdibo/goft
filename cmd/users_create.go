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
	return &cobra.Command{
		Use:   "create email first_name last_name kind campus_id",
		Short: "Create a new user",
		Long: `Create a new user account.
This command requires the Advanced tutor role
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
			_ = (*api).CreateUser(&user)
			return nil
		},
	}
}
var userCreateCmd = NewUserCreateCmd(&API)

func init() {
	usersCmd.AddCommand(userCreateCmd)

	// login - password

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userCreateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userCreateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	userCreateCmd.Flags().String("login", "", "Set a custom login, leave empty to let the intra generate one")
}
