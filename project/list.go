package project

import (
	"fmt"
	"log"
	"os"

	"github.com/admdwrf/ovhcli/internal"
	"github.com/spf13/cobra"
)

var cmdProjectList = &cobra.Command{
	Use:   "list",
	Short: "List all projects: ovhcli project list",
	Run: func(cmd *cobra.Command, args []string) {
		c := internal.Client()
		if c == nil {
			os.Exit(1)
		}
		projects := internal.Projects{}
		if err := c.Get("/cloud/project", &projects); err != nil {
			log.Fatal(err)
		}
		for _, p := range projects {
			fmt.Printf("project: %s\n", p)
		}
	},
}
