package cmd

import (
	"bytes"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestNewRootCmd(t *testing.T) {
	b := bytes.NewBufferString("")
	rootCmd := NewRootCmd()
	assert.IsType(t, &cobra.Command{}, rootCmd)
	rootCmd.SetOut(b)
	err := rootCmd.Execute()
	assert.Nil(t, err)
	out, err := ioutil.ReadAll(b)
	assert.Equal(t, "CLI tool to interact with 42's API\n\n", string(out))
}