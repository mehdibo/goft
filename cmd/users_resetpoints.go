package cmd

import (
	"errors"
	"goft/pkg/ftapi"
	"strconv"

	"github.com/spf13/cobra"
)

// NewResetPointsCmd create reset points command
func NewResetPointsCmd(api *ftapi.APIInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "reset-points login points reason",
		Short: "Reset correction points for a user",
		Long: "This command requires the Advanced tutor role",
		Args: func(cmd *cobra.Command, args []string) error {
			if err := cobra.ExactArgs(3)(cmd, args); err != nil {
				return err
			}
			points, err := strconv.ParseUint(args[1], 10, 32)
			if err != nil || points < 0 {
				return errors.New("points must be greater than or equal to 0")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			// Get current user points
			user, err := (*api).GetUserByLogin(args[0])
			if err != nil {
				return err
			}
			// If not zero, delete all points
			if user.CorrectionPoints > 0 {
				err = (*api).RemoveCorrectionPoints(user.Login, uint(user.CorrectionPoints), args[2])
				if err != nil {
					return err
				}
			}
			// If less than zero reset it to zer
			if user.CorrectionPoints < 0 {
				err = (*api).AddCorrectionPoints(user.Login, uint(user.CorrectionPoints * -1), args[2])
				if err != nil {
					return err
				}
			}
			// Add required points
			points, _ := strconv.ParseUint(args[1], 10, 0)
			if points == 0 {
				return nil
			}
			return (*api).AddCorrectionPoints(user.Login, uint(points), args[2])
		},
	}
}

var resetPointsCmd = NewResetPointsCmd(&API)

func init() {
	usersCmd.AddCommand(resetPointsCmd)
}