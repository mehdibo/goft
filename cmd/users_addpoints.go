package cmd

import (
	"errors"
	"goft/pkg/ftapi"
	"strconv"

	"github.com/spf13/cobra"
)

// NewAddPointsCmd create the add points cmd
func NewAddPointsCmd(api *ftapi.APIInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "add-points login points reason",
		Short: "Add correction points to user",
		Long: "This command requires the Advanced tutor role",
		Args: func(cmd *cobra.Command, args []string) error {
			if err := cobra.ExactArgs(3)(cmd, args); err != nil {
				return err
			}
			points, err := strconv.ParseUint(args[1], 10, 32)
			if err != nil || points <= 0 {
				return errors.New("points must be greater than 0")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			points, _ := strconv.ParseUint(args[1], 10, 0)
			return (*api).AddCorrectionPoints(args[0], uint(points), args[2])
		},
	}
}

// addPointsCmd represents the addPoints command
var addPointsCmd = NewAddPointsCmd(&API)

func init() {
	usersCmd.AddCommand(addPointsCmd)
}
