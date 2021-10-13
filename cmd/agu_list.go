package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"goft/pkg/ftapi"
)

// NewAguListCmd Create the agu list cmd
func NewAguListCmd(api *ftapi.APIInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "list login",
		Short: "List a user's AGUs",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			agus, err := (*api).GetUserAgus(args[0])
			if err != nil {
				return err
			}
			for _, agu := range agus {
				_, _ = fmt.Fprint(cmd.OutOrStdout(), "\n")
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "ID: %d\n", agu.ID)
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Reason: %s\n", agu.Reason)
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Is free: %t\n", agu.IsFree)
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Begins at: %s\n", agu.BeginDate)
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Expected end: %s\n", agu.ExpectedEndDate)
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Ends at: %s\n", agu.EndDate)
				_, _ = fmt.Fprintf(cmd.OutOrStdout(), "Created at: %s\n", agu.CreatedAt)
			}
			return nil
		},
	}
}
var aguListCmd = NewAguListCmd(&API)

func init() {
	aguCmd.AddCommand(aguListCmd)
}
