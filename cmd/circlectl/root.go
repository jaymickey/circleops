/*
Copyright © 2019 Josh Michielsen <github@mickey.dev>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile       string
	noInteractive bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "circlectl",
	Short: "CLI tool for managing workflows and jobs within CircleCI",
	Long: `The official CircleCI CLI lacks features, and has a clunky
	interface. This CLI focuses on interactions with the workflow,
	and jobs aspects of CircleCI - such as allowing developers to
	easily re-run a job or workflow.`,
	//Run: func(cmd *cobra.Command, args []string) {
	//fmt.Printf("Using config file: %v", viper.ConfigFileUsed())

	//},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVar(&noInteractive, "no-interactive", false, "")
	_ = rootCmd.PersistentFlags().MarkHidden("no-interactive")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.circlectl/config.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	viper.SetDefault("serverURL", "https://circleci.deployment-dev.cni.digital")
	viper.SetDefault("apiLocation", path.Join(viper.GetString("serverURL"), "api", "v1.1"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(path.Join(home, ".circlectl"))
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func ensureConfigPath() (configDir string, err error) {
	configDir = filepath.Dir(viper.ConfigFileUsed())

	if viper.ConfigFileUsed() == "" {
		home, err := homedir.Dir()
		if err != nil {
			return "", fmt.Errorf("error getting home dir: %v", err)
		}
		configDir = path.Join(home, ".circlectl")
	}

	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
			return "", fmt.Errorf("error creating directory; %v: %v", configDir, err)
		}
	}

	return configDir, nil
}
