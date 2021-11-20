package cmd

import (
	"goft/pkg/ftapi"
	"log"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func NewGetProjectListCmd(api *ftapi.APIInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "Registered project list",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			defer saveConfig()
			me, err := (*api).GetMe()
			if err != nil {
				return err
			}
			id := me.ID
			filter := map[string]string{
				"marked": "false",
			}
			projects, err := (*api).GetUserProjects(id, filter, nil)
			if err != nil {
				color.Set(color.FgRed)
				log.Print("GetUserProjects:", err)
				color.Set(color.Reset)
				return err
			}
			for _, project := range projects {
				//				color.Set(color.FgGreen)
				cmd.Println(project.Project.Slug)
				/*
					color.Set(color.Reset)
					if len(project.Teams) != 0 {
						cmd.Println(project.Teams[0].RepoURL)
					}
				*/
			}
			return nil
		},
	}
}

var getProjectListCmd = NewGetProjectListCmd(&API)

func init() {
	projectsCmd.AddCommand(getProjectListCmd)
}
