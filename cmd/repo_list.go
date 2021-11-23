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
			var isQuiet bool
			if isQuiet, err = cmd.PersistentFlags().GetBool("quiet"); err != nil {
				return err
			}

			if !isQuiet {
				fmt.Printf("\nShowing %d projects in @%s\n\n", limit, user)
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
					if len(project.Teams) == 0 {
						continue
					}
					count++

					team, err := currentTeam(project)
					if err != nil {
						return err
					}

					var repo string
					if team.RepoURL == "" {
						repo = "repository not found"
					} else {
						repo = team.RepoURL
					}

					if isQuiet {
						fmt.Println(project.Project.Slug)
					} else {
						fmt.Printf("%-25s %s\n", project.Project.Slug, repo)
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
	getProjectListCmd.PersistentFlags().BoolP("quiet", "q", false, "Only display project slug")
	projectsCmd.AddCommand(getProjectListCmd)
}
