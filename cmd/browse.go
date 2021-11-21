package cmd

import (
	"errors"
	"fmt"
	"goft/pkg/ftapi"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/spf13/cobra"
)

// browseCmd represents the browse command
func NewBrowseCmd(api *ftapi.APIInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "browse <command>",
		Short: "Open the project page in the web browser",
		Long:  `Open the project page in the web browser`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			defer saveConfig()
			url, err := getProjectURL(api, args[0])
			if err != nil {
				return err
			}
			return execBrowser(url)
		},
	}
}

func getProjectURL(api *ftapi.APIInterface, slug string) (string, error) {
	var url string

	user, err := (*api).GetMe()
	if err != nil {
		return "", err
	}

	if is42Cursus(api, slug) {
		baseURL := "https://projects.intra.42.fr/"
		url = baseURL + slug + "/" + user.Login
		return url, nil
	} else {
		baseURL := "https://projects.intra.42.fr/projects/"
		id := user.ID
		for i := 1; ; i++ {
			projects, err := (*api).GetUserProjects(id, nil, nil, i)
			if err != nil {
				return "", err
			}
			if len(projects) == 0 {
				break
			}
			for _, project := range projects {
				if slug == project.Project.Slug {
					url = baseURL + slug + "/projects_users/" + strconv.Itoa(project.ID)
					return url, nil
				}
			}
		}
	}
	return "", errors.New("project not found")
}

func is42Cursus(api *ftapi.APIInterface, slug string) bool {
	project, err := (*api).GetProjectByName(slug)
	if err != nil {
		return false
	}
	cursuses := project.Cursus
	for _, cursus := range cursuses {
		if cursus.Slug == "42cursus" {
			return true
		}
	}
	return false
}

func execBrowser(url string) error {
	browser := "xdg-open"
	if runtime.GOOS == "darwin" {
		browser = "open"
	}
	args := []string{url}
	browser, err := exec.LookPath(browser)
	if err == nil {
		cmd := exec.Command(browser, args...)
		cmd.Stderr = os.Stderr
		err = cmd.Start()
		if err != nil {
			return fmt.Errorf("cannot start command: %v", err)
		}
	} else {
		fmt.Println("Open this URL")
		fmt.Println(url)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(NewBrowseCmd(&API))
}
