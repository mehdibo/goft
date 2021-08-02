package cmd

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"goft/pkg/ftapi"
	"io/ioutil"
)

// NewRequestsPostCmd Create the post request cmd
func NewRequestsPostCmd(api *ftapi.APIInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "post path",
		Short: "Send a POST request",
		Long: `Send a POST request to path, the path must be the part after /v2/`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Print("Enter payload: (press CTRL+D when you finish)\n")
			var payload string
			scanner := bufio.NewScanner(cmd.InOrStdin())
			for scanner.Scan() {
				line := scanner.Text()
				payload += line+"\n"
			}
			fmt.Println(payload)
			resp, err := (*api).Post(args[0], "application/json", bytes.NewReader([]byte(payload)))
			if err != nil {
				return err
			}
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			var prettyJson bytes.Buffer
			_ = json.Indent(&prettyJson, bodyBytes, "", "\t")
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%s\n", resp.Status)
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%s\n", prettyJson.String())
			return nil
		},
	}
}
var requestsPostCmd = NewRequestsPostCmd(&API)

func init() {
	requestsCmd.AddCommand(requestsPostCmd)
}
