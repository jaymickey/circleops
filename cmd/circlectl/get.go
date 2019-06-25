package main

import "github.com/spf13/cobra"

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get one, many, or all of a particular resouce.",
	Long: `Display one or many of a desired resource.

This command will print a table of the most important information about the specified
resource(s).

Examples:

	# List projects
	circlectl get projects`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
