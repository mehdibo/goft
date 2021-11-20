package cmd

import (
	"errors"
	"fmt"
	"goft/pkg/ftapi"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func NewCloneProjectCmd(api *ftapi.APIInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "clone",
		Short: "Clone project repogitory",
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			defer saveConfig()
			me, err := (*api).GetMe()
			if err != nil {
				return err
			}
			id := me.ID
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
					color.Set(color.Reset)
					if args[0] != project.Project.Slug {
						continue
					}
					if len(project.Teams) == 0 {
						continue
					}
					var targetPath string
					if len(args) == 2 {
						targetPath = args[1]
					} else {
						targetPath = args[0]
					}
					team, err := currentTeam(project)
					if err != nil {
						return err
					}
					return cloneRepo(team.RepoURL, targetPath)
				}
			}
			cmd.Printf("Team of %s is not locked.", args[0])
			return fmt.Errorf("%s is not in your projects", args[0])
		},
	}
}

func currentTeam(project *ftapi.ProjectUser) (*ftapi.Team, error) {
	for _, team := range project.Teams {
		if project.CurrentTeamID == team.ID {
			return &team, nil
		}
	}
	return nil, errors.New("not found")
}

func cloneRepo(repoURL, targetPath string) error {
	git, err := exec.LookPath("git")
	if err != nil {
		return err
	}
	var args []string
	if targetPath != "" {
		args = []string{"clone", repoURL, targetPath}
	} else {
		args = []string{"clone", repoURL}
	}
	cmd := exec.Command(git, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("cannot exec git: %v", err)
	}
	return nil
}

var cloneProjectCmd = NewCloneProjectCmd(&API)

func init() {
	projectsCmd.AddCommand(cloneProjectCmd)
}
