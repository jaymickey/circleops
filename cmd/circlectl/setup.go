package main

import (
	"bufio"
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Setup URL and API token",
	RunE:  setup,
}

var stdin *os.File

func setup(cmd *cobra.Command, args []string) error {
	if stdin == nil {
		stdin = os.Stdin
	}
	sc := bufio.NewScanner(stdin)

	fmt.Print("CircleCI Base URL (e.g. https://circleci.com): ")
	sc.Scan()
	url := sc.Text()

	fmt.Print("API Token: ")
	sc.Scan()
	api := sc.Text()

	viper.Set("serverURL", url)
	viper.Set("apiToken", api)
	viper.Set("apiLocation", path.Join(viper.GetString("serverURL"), "api", "v1.1"))

	if !noInteractive {
		configDir, err := ensureConfigPath()
		if err != nil {
			return fmt.Errorf("unable to ensure config path exists: %v", err)
		}
		if err := viper.WriteConfigAs(path.Join(configDir, "config.yaml")); err != nil {
			return fmt.Errorf("couldn't save config to file: %v", err)
		}
	}

	return nil
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
