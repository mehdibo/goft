package cmd

import (
	"errors"
	"fmt"
	"goft/pkg/ftapi"

	"github.com/spf13/cobra"
)

func formatProjectText(project *ftapi.Project) string {
	output := fmt.Sprintf(`ID: %d
Name: %s
Description: %s
Cursus: %s
Objectives: %s
Estimate time: %s
`,
		project.ID,
		project.Name,
		project.ProjectSessions[0].Description,
		project.Cursus[0].Name,
		project.ProjectSessions[0].Objectives,
		project.ProjectSessions[0].EstimateTime,
	)
	return output
}

func NewGetProjectCmd(api *ftapi.APIInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Short: "Get details about a project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			project, err := (*api).GetProjectByName(args[0])
			if err != nil {
				return err
			}
			if project == nil {
				return errors.New("failed getting project")
			}
			cmd.Print(formatProjectText(project))
			return nil
		},
	}
}

var getProjectCmd = NewGetProjectCmd(&API)

func init() {
	projectsCmd.AddCommand(getProjectCmd)
}
