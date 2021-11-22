package cmd

import (
	"fmt"
	"goft/pkg/ftapi"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

// browseCmd represents the browse command
func NewBrowseCmd(api *ftapi.APIInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "browse <project>",
		Short: "Open the project page in the web browser",
		Long:  `Open the project page in the web browser`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			defer saveConfig()
			baseURL := "https://projects.intra.42.fr/"
			url := baseURL + args[0] + "/mine"
			return execBrowser(url)
		},
	}
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
