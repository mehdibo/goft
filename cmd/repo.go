package cmd

import (
	"github.com/spf13/cobra"
)

// NewUsersCmd Creates new users command
func NewProjectsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "repo",
		Short: "Handle your registerd repogitory",
	}

}

var projectsCmd = NewProjectsCmd()

func init() {
	rootCmd.AddCommand(projectsCmd)
}
