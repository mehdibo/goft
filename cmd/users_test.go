package cmd

import (
	"bytes"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUsersCmd(t *testing.T) {
	b := bytes.NewBufferString("")
	usersCmd := NewUsersCmd()
	assert.IsType(t, &cobra.Command{}, usersCmd)
	usersCmd.SetOut(b)
	err := usersCmd.Execute()
	assert.Nil(t, err)
}