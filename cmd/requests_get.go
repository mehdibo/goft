package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"goft/pkg/ftapi"
	"io/ioutil"
)

// NewRequestsGetCmd Create the close create cmd
func NewRequestsGetCmd(api *ftapi.APIInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "get path",
		Short: "Send a GET request",
		Long: `Send a GET request to path, the path must be the part after /v2/`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*api).Get(args[0])
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
var requestsGetCmd = NewRequestsGetCmd(&API)

func init() {
	requestsCmd.AddCommand(requestsGetCmd)
}
