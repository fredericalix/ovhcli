package caas

import (
	"fmt"
	"os"

	"github.com/admdwrf/ovhcli/sdk"
	"github.com/spf13/cobra"
)

var containerservid string

func init() {
	cmdContainersServiceInfo.PersistentFlags().StringVarP(&containerservid, "serviceName", "", "", "Containers Service Name")
}

var cmdContainersServiceInfo = &cobra.Command{
	Use:   "info",
	Short: "Info about a project: ovhcli caas info",
	Run: func(cmd *cobra.Command, args []string) {

		c, err := sdk.NewClient()
		if err != nil {
			fmt.Printf("Error: %s", err)
			os.Exit(1)
		}

		containersserviceinfo, err := c.ContainersServiceInfo(containerservid)
		if err != nil {
			fmt.Printf("Error: %s", err)
			os.Exit(1)
		}
		fmt.Printf("Service Name: %s\n", containersserviceinfo.Name)
		fmt.Printf("Cluster: %s\n", containersserviceinfo.Cluster)
		fmt.Printf("Load Balancer name: %s\n", containersserviceinfo.LoadBalancer)
		fmt.Printf("Creation Date: %s\n", containersserviceinfo.CreationDate)
	},
}
