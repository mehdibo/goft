package cmd

import (
	"github.com/spf13/cobra"
)

// NewUsersCmd Creates new users command
func NewProjectsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "projects",
		Short: "Interact with the project entity",
	}

}

var projectsCmd = NewProjectsCmd()

func init() {
	rootCmd.AddCommand(projectsCmd)
}
