package project

import (
	"fmt"
	"os"

	"github.com/admdwrf/ovhcli/sdk/cloud"
	"github.com/spf13/cobra"
)

var cmdProjectList = &cobra.Command{
	Use:   "list",
	Short: "List all projects: ovhcli project list",
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := cloud.ProjectList()
		if err != nil {
			fmt.Printf("Error: %s", err)
			os.Exit(1)
		}
		for _, p := range projects {
			fmt.Printf("%s\n", p)
		}
	},
}
