package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"goft/pkg/ftapi"
)

func isValidKind(kind string) bool {
	switch kind {
	case "agu",
		"black_hole",
		"deserter",
		"non_admitted",
		"serious_misconduct",
		"social_security",
		"other":
		return true
	}
	return false
}

// NewCloseCreateCmd Create the close create cmd
func NewCloseCreateCmd(api *ftapi.APIInterface) *cobra.Command {
	// login - kind - reason -
	return &cobra.Command{
		Use:   "create login kind reason closer_login",
		Short: "Create a new close for a user",
		Long: `This command requires the Basic staff role

kind must be one of the following options: agu, black_hole, deserter, non_admitted, serious_misconduct, social_security or other`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 4 {
				return fmt.Errorf("accepts 3 arg(s), received %d", len(args))
			}
			if !isValidKind(args[1]) {
				return fmt.Errorf("'%s' is not a valid kind", args[1])
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			closer, err := (*api).GetUserByLogin(args[3])
			if err != nil {
				return err
			}
			if closer == nil || closer.ID == 0 {
				return errors.New("couldn't get closer")
			}
			err = (*api).CreateClose(&ftapi.Close{
				Kind: args[1],
				Reason: args[2],
				CommunityServices: nil,
				User: &ftapi.User{
					Login: args[0],
				},
				Closer: &ftapi.User{
					ID: closer.ID,
				},
			})
			return err
		},
	}
}
var closesCreateCmd = NewCloseCreateCmd(&API)

func init() {
	closesCmd.AddCommand(closesCreateCmd)
}
