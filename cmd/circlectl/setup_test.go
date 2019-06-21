package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestSetupCmd(t *testing.T) {
	expURL := "https://circleci.com"
	expToken := "test_key"
	input := expURL + "\n" + expToken
	tmp, _ := ioutil.TempFile(os.TempDir(), "stdin")
	defer os.Remove(tmp.Name())

	_, err := tmp.WriteString(input)
	if err != nil {
		t.Fatalf("error writing string to tmp file: %v", err)
	}

	_, err = tmp.Seek(0, 0)
	if err != nil {
		t.Fatalf("error seeking to start of tmp file: %v", err)
	}

	stdin = tmp

	setupCmdArgs := []string{
		"setup",
		"--no-interactive",
	}

	_, err = executeCommandC(rootCmd, setupCmdArgs...)
	assert.NoError(t, err, "command threw an error")
	assert.Equalf(t, expURL, viper.GetString("serverURL"), "server url did not match expected value. expected: \"%v\" got: \"%v\"", expURL, viper.GetString("serverURL"))
	assert.Equalf(t, "test_key", viper.GetString("apiToken"), "api token did not match expected value. expected: \"%v\" got: \"%v\"", expToken, viper.GetString("apiToken"))
}
