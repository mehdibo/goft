package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClosesCmd(t *testing.T) {
	closesCmd := NewClosesCmd()
	assert.IsType(t, &cobra.Command{}, closesCmd)
}