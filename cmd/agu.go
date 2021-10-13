package cmd

import (
	"github.com/spf13/cobra"
)

// NewAguCmd creates agu cmd
func NewAguCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "agu",
		Short: "Manage Anti Grav Units (AGUs)",
	}
}
var aguCmd = NewAguCmd()

func init() {
	rootCmd.AddCommand(aguCmd)
}
