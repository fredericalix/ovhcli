package cloud

import (
	"fmt"
	"os"

	"github.com/admdwrf/ovhcli/sdk"
	"github.com/spf13/cobra"
)

var cmdProjectList = &cobra.Command{
	Use:   "list",
	Short: "List all projects: ovhcli project list",
	Run: func(cmd *cobra.Command, args []string) {

		c, err := sdk.NewClient()
		if err != nil {
			fmt.Printf("Error: %s", err)
			os.Exit(1)
		}

		projects, err := c.CloudProjectList()
		if err != nil {
			fmt.Printf("Error: %s", err)
			os.Exit(1)
		}
		for _, p := range projects {
			fmt.Printf("%s\n", p)
		}
	},
}
