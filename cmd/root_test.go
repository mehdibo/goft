package cmd

import (
	"bytes"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestNewRootCmd(t *testing.T) {
	viper.Set("client_id", "test_id")
	viper.Set("client_secret", "test_secret")
	b := bytes.NewBufferString("")
	rootCmd := NewRootCmd()
	assert.IsType(t, &cobra.Command{}, rootCmd)
	rootCmd.SetOut(b)
	err := rootCmd.Execute()
	assert.Nil(t, err)
	out, _ := ioutil.ReadAll(b)
	// Command outputs nothing as it doesnt load a config file
	// And required configs are set using viper directly
	assert.Equal(t, "", string(out))
}
