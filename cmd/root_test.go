package cmd

import (
	"bytes"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
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
}
