package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/jaymickey/circleops/api"
	"github.com/jaymickey/circleops/client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"mickey.dev/circleops/api"
	"mickey.dev/circleops/client"
)

var getProjectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Get all projects the user has followed.",
	Run:   getProjects,
}

func init() {
	getCmd.AddCommand(getProjectsCmd)
}

func getProjects(cmd *cobra.Command, args []string) {
	host, port, token := getParts()
	endpoint := "projects"
	projectsLogger := log.WithFields(log.Fields{
		"host":     host,
		"endpoint": endpoint,
		"port":     port,
	})

	cl, err := client.NewClient(host, port, endpoint, token)
	if err != nil {
		projectsLogger.Fatalf("error creating client: %v", err)
	}

	projects, err := api.GetProjects(cl)
	if err != nil {
		projectsLogger.Fatalf("error retrieving projects from API: %v", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 24, 8, 0, ' ', 0)
	fmt.Fprintln(w, "Name\tOrganisation\tLanguage")
	fmt.Fprintln(w, "--------\t--------\t--------")

	for _, proj := range projects {
		fmt.Fprintf(w, "%s\t%s\t%s\n", proj.Name, proj.Organisation, proj.Language)
	}

	w.Flush()
}
