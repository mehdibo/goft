package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"goft/pkg/ftapi"
	"os"
)

// NewSetImgCmd create the users setimg cmd
func NewSetImgCmd(api *ftapi.APIInterface) *cobra.Command {
	cmd := cobra.Command{
		Use:   "setimg login image_path",
		Short: "Set image for user",
		Long: `This command requires the Advanced tutor role.
Image file must be 3Kb and 1Mb
`,
		Args: func(cmd *cobra.Command, args []string) error {
			if err := cobra.ExactArgs(2)(cmd, args); err != nil {
				return err
			}
			const minImgLen = 3072
			const maxImgLen = 1048576
			// Check if file exists
			imgFile, err := os.Stat(args[1])
			if err != nil {
				return err
			}
			if imgFile.Size() < minImgLen || imgFile.Size() > maxImgLen {
				return errors.New("image file size is invalid")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			imgFile, err := os.Open(args[1])
			if err != nil {
				return err
			}
			defer imgFile.Close()
			err = (*api).SetUserImage(args[0], imgFile)
			if err != nil {
				return err
			}
			return nil
		},
	}
	return &cmd
}
var setimgCmd = NewSetImgCmd(&API)

func init() {
	usersCmd.AddCommand(setimgCmd)
}