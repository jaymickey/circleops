package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"mickey.dev/circleops/api"
	"mickey.dev/circleops/client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	includeUser bool
)

var getOrganisationsCmd = &cobra.Command{
	Use:     "organisations [--include-user]",
	Aliases: []string{"org", "orgs"},
	Run:     getOrganisations,
}

func init() {
	getCmd.AddCommand(getOrganisationsCmd)
	getOrganisationsCmd.Flags().BoolVar(&includeUser, "include-user", false, "include the authenticated user in the returned list.")
}

func getOrganisations(cmd *cobra.Command, args []string) {
	host, port, token := getParts()
	endpoint := "user/organizations"

	orgsLogger := log.WithFields(log.Fields{
		"host":     host,
		"endpoint": endpoint,
		"port":     port,
	})

	cl, err := client.NewClient(host, port, endpoint, token)
	if err != nil {
		orgsLogger.Fatalf("error creating client: %v", err)
	}

	orgs, err := api.GetOrganisations(cl, includeUser)
	if err != nil {
		orgsLogger.Fatalf("error retrieving organisations from API: %v", err)
	}

	w := tabwriter.NewWriter(os.Stdout, 24, 8, 4, ' ', 0)
	fmt.Fprintln(w, "Name\tAdmin\tVCS Provider")
	fmt.Fprintln(w, "--------\t--------\t--------")

	for _, org := range orgs {
		fmt.Fprintf(w, "%s\t%v\t%s\n", org.Name, org.IsAdmin, org.VCS)
	}
	w.Flush()
}
