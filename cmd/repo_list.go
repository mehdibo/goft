package cmd

import (
	"goft/pkg/ftapi"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func NewGetProjectListCmd(api *ftapi.APIInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "Show team locked project list",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			defer saveConfig()
			var limit int
			var err error
			if limit, err = cmd.PersistentFlags().GetInt("limit"); err != nil {
				return err
			}
			me, err := (*api).GetMe()
			if err != nil {
				return err
			}
			id := me.ID
			count := 0
		loop:
			for i := 1; ; i++ {
				projects, err := (*api).GetUserProjects(id, nil, nil, i)
				if err != nil {
					color.Set(color.FgRed)
					cmd.PrintErr("GetUserProjects:", err)
					color.Set(color.Reset)
					return err
				}
				if len(projects) == 0 {
					break
				}
				for _, project := range projects {
					if len(project.Teams) != 0 {
						cmd.Println(project.Project.Slug)
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
	projectsCmd.AddCommand(getProjectListCmd)
}
