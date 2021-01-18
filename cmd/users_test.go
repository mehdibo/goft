package cmd

import (
	"bytes"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestNewUsersCmd(t *testing.T) {
	viper.Set("client_id", "test_id")
	viper.Set("client_secret", "test_secret")
	b := bytes.NewBufferString("")
	usersCmd := NewUsersCmd()
	assert.IsType(t, &cobra.Command{}, usersCmd)
	usersCmd.SetOut(b)
	err := usersCmd.Execute()
	assert.Nil(t, err)
	out, _ := ioutil.ReadAll(b)
	// Command outputs nothing as it doesnt load a config file
	// And required configs are set using viper directly
	assert.Equal(t, "", string(out))
}