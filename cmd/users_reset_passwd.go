package cmd

import (
	"errors"
	"fmt"
	"github.com/sethvargo/go-password/password"
	"github.com/spf13/cobra"
	"goft/pkg/ftapi"
	"gopkg.in/gomail.v2"
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
			smtpUser, _ := cmd.LocalFlags().GetString("smtp-user")
			smtpPass, _ := cmd.LocalFlags().GetString("smtp-pass")
			smtpHost, _ := cmd.LocalFlags().GetString("smtp-host")
			smtpPort, _ := cmd.LocalFlags().GetInt("smtp-port")
			fromEmail, _ := cmd.LocalFlags().GetString("from-email")
			if smtpUser == "" || smtpPass == "" || smtpHost == "" || smtpPort <= 0 || fromEmail == "" {
				return errors.New("missing required flags")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			smtpUser, _ := cmd.LocalFlags().GetString("smtp-user")
			smtpPass, _ := cmd.LocalFlags().GetString("smtp-pass")
			smtpHost, _ := cmd.LocalFlags().GetString("smtp-host")
			smtpPort, _ := cmd.LocalFlags().GetInt("smtp-port")
			fromEmail, _ := cmd.LocalFlags().GetString("from-email")
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
			err = (*api).UpdateUser(user.Login, &newUser)
			if err != nil {
				return err
			}
			// Send the password via email
			mailDialer := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
			m := gomail.NewMessage()
			m.SetHeader("From", fromEmail)
			m.SetAddressHeader("To", user.Email, user.FirstName)
			m.SetHeader("Subject", "New password for your 42 Account")

			htmlBody := fmt.Sprintf(`<html>
Hello <b>%s</b>!<br />
<br />
This is your new 42 Intranet password: <b>%s</b><br />
<br />
You can use it to login <a href="https://signin.intra.42.fr/users/sign_in">here</a><br />
<br />
<b>Do not reply to this email, if you still have a problem use Slack to report it.</b><br />
</html>
`, user.FirstName, newPass)
			txtBody := fmt.Sprintf(`Hello %s!

This is your new 42 Intranet password: %s

You can use it to login here: https://signin.intra.42.fr/users/sign_in

Do not reply to this email, if you still have a problem use Slack to report it.
`, user.FirstName, newPass)
			m.SetBody("text/html", htmlBody)
			m.AddAlternative("text/plain", txtBody)

			err = mailDialer.DialAndSend(m)
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().String("from-email", "", "Address to be used as a sender")
	cmd.Flags().String("smtp-user", "", "SMTP username")
	cmd.Flags().String("smtp-pass", "", "SMTP password")
	cmd.Flags().String("smtp-host", "", "SMTP host")
	cmd.Flags().Int("smtp-port", 25, "SMTP port")
	_ = cmd.MarkFlagRequired("from-email")
	_ = cmd.MarkFlagRequired("smtp-user")
	_ = cmd.MarkFlagRequired("smtp-pass")
	_ = cmd.MarkFlagRequired("smtp-host")
	_ = cmd.MarkFlagRequired("smtp-port")
	return &cmd
}

var passwdGenerator, _ = password.NewGenerator(nil)
var resetPasswdCmd = NewResetPasswdCmd(&API, passwdGenerator)

func init() {
	usersCmd.AddCommand(resetPasswdCmd)
}
