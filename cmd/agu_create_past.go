package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"goft/pkg/ftapi"
	"strconv"
)

// NewAguCreatePastCmd Create the agu create_past cmd
func NewAguCreatePastCmd(api *ftapi.APIInterface) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create_past login duration",
		Short: "List a user's AGUs",
		Long: "Add a free AGU in the past for a user to delay the blackhole, duration must be in days",
		Args: func(cmd *cobra.Command, args []string) error {
			n := 2
			if len(args) != n {
				return fmt.Errorf("accepts %d arg(s), received %d", n, len(args))
			}
			duration, err := strconv.ParseInt(args[1], 10, 32)
			if err != nil || duration <= 0 {
				return errors.New("duration must be greater than 0")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			duration, err := strconv.ParseInt(args[1], 10, 32)
			if err != nil {
				return err
			}
			reason, err := cmd.Flags().GetString("reason")
			if err != nil {
				return err
			}
			err = (*api).CreateFreePastAgu(args[0], int(duration), reason)
			if err != nil {
				return err
			}
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "A free past AGU was created of %d days for %s", duration, args[0])
			return nil
		},
	}
	cmd.Flags().String("reason", "", "Reason for the freeze")
	return cmd
}
var aguCreatePastCmd = NewAguCreatePastCmd(&API)

func init() {
	aguCmd.AddCommand(aguCreatePastCmd)
}
