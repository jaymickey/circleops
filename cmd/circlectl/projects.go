package main

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/jaymickey/circleops/api"
	"github.com/jaymickey/circleops/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Get all projects the user has followed.",
	Run:   getProjects,
}

func init() {
	getCmd.AddCommand(projectsCmd)
}

func getProjects(cmd *cobra.Command, args []string) {
	host := viper.GetString("apilocation")
	port := viper.GetString("port")
	endpoint := "projects"
	token := viper.GetString("apiToken")

	cl, err := client.NewClient(host, endpoint, port, token)
	if err != nil {
		log.Fatalf("error creating client: %v", err)
	}

	projects := api.GetProjects(cl)

	w := tabwriter.NewWriter(os.Stdout, 24, 8, 0, ' ', 0)
	fmt.Fprintln(w, "Name\tOrganisation\tLanguage")
	fmt.Fprintln(w, "--------\t--------\t--------")

	for _, proj := range projects {
		fmt.Fprintf(w, "%s\t%s\t%s\n", proj.Name, proj.Organisation, proj.Language)
	}

	w.Flush()
}
