package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/spf13/cobra"
)

func executeCommandC(root *cobra.Command, args ...string) (output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOutput(buf)
	root.SetArgs(args)
	err = root.Execute()
	return buf.String(), err
}

func TestRootCommand(t *testing.T) {
	rootCmdArgs := []string{
		"--help",
	}

	output, err := executeCommandC(rootCmd, rootCmdArgs...)
	assert.NoError(t, err, "command threw an error")
	assert.NotEqual(t, "", output, "output was nil")
}
