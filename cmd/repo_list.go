package cmd

import (
	"fmt"
	"goft/pkg/ftapi"
	"os"

	"github.com/spf13/cobra"
)

func NewGetProjectListCmd(api *ftapi.APIInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "Show team locked project list",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			var limit int
			var err error
			if limit, err = cmd.PersistentFlags().GetInt("limit"); err != nil {
				return err
			}
			var user string
			if user, err = cmd.PersistentFlags().GetString("user"); err != nil {
				return err
			}
			count := 0
		loop:
			for i := 1; ; i++ {
				projects, err := (*api).GetUserProjects(user, nil, nil, i)
				if err != nil {
					return err
				}
				if len(projects) == 0 {
					break
				}
				for _, project := range projects {
					if len(project.Teams) != 0 {
						fmt.Println(project.Project.Slug)
						count++
					}
					if count >= limit {
						break loop
					}
				}
			}
			return nil
		},
	}
}

var getProjectListCmd = NewGetProjectListCmd(&API)

func init() {
	getProjectListCmd.PersistentFlags().IntP("limit", "L", 5, "Maximum number of repositories to list")
	getProjectListCmd.PersistentFlags().StringP("user", "u", os.Getenv("USER"), "Set specific user")
	projectsCmd.AddCommand(getProjectListCmd)
}
