package caas

import (
	"fmt"
	"os"

	"github.com/admdwrf/ovhcli/sdk"
	"github.com/spf13/cobra"
)

var cmdContainersServicesList = &cobra.Command{
	Use:   "list",
	Short: "List all containers services: ovhcli caas list",
	Run: func(cmd *cobra.Command, args []string) {

		c, err := sdk.NewClient()
		if err != nil {
			fmt.Printf("Error: %s", err)
			os.Exit(1)
		}

		containersservices, err := c.ContainersServicesList()
		if err != nil {
			fmt.Printf("Error: %s", err)
			os.Exit(1)
		}
		for _, p := range containersservices {
			fmt.Printf("%s\n", p)
		}
	},
}
