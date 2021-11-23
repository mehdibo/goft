package cmd

import (
	"errors"
	"fmt"
	"goft/pkg/ftapi"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func NewCloneProjectCmd(api *ftapi.APIInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "clone <project slug>",
		Short: "Clone a repogitory locally",
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			user, err := cmd.PersistentFlags().GetString("user")
			if err != nil {
				return err
			}
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
					if args[0] != project.Project.Slug {
						continue
					}
					if len(project.Teams) == 0 {
						break loop
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
					if team.RepoURL == "" {
						fmt.Fprintf(os.Stderr, "repository not found: %s\n", args[0])
						return nil
					}
					return cloneRepo(team.RepoURL, targetPath)
				}
			}
			return fmt.Errorf("%s's team is not locked.", args[0])
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
	cloneProjectCmd.PersistentFlags().StringP("user", "u", os.Getenv("USER"), "Set specific user")
	projectsCmd.AddCommand(cloneProjectCmd)
}
