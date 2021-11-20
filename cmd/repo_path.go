package cmd

import (
	"fmt"
	"goft/pkg/ftapi"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func NewRepoPathCmd(api *ftapi.APIInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "path",
		Short: "Show project repogitory path",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			defer saveConfig()
			me, err := (*api).GetMe()
			if err != nil {
				return err
			}
			id := me.ID
			projects, err := (*api).GetUserProjects(id, nil, nil)
			if err != nil {
				color.Set(color.FgRed)
				cmd.PrintErr("GetUserProjects:", err)
				color.Set(color.Reset)
				return err
			}
			for _, project := range projects {
				color.Set(color.Reset)
				if args[0] != project.Project.Slug {
					continue
				}
				if len(project.Teams) == 0 {
					continue
				}
				team := latestTeam(project.Teams)
				fmt.Println(team.RepoURL)
				return nil
			}
			cmd.Printf("Team of %s is not locked.", args[0])
			return fmt.Errorf("%s is not in your projects", args[0])
		},
	}
}

var repoPathCmd = NewRepoPathCmd(&API)

func init() {
	projectsCmd.AddCommand(repoPathCmd)
}
