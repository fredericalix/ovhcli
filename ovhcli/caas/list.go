package caas

import (
	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

	"github.com/spf13/cobra"
)

var withDetails bool

func init() {
	cmdContainersServicesList.PersistentFlags().BoolVarP(&withDetails, "withDetails", "", false, "Display containers details")
}

var cmdContainersServicesList = &cobra.Command{
	Use:   "list",
	Short: "List all containers services: ovhcli caas list",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ovh.NewClient()
		common.Check(err)

		containersservices, err := client.ContainersServicesList()
		common.Check(err)

		if withDetails {
			contComplete := []ovh.ContainersService{}
			for _, cont := range containersservices {
				c, err := client.ContainersServiceInfo(cont.Name)
				common.Check(err)
				contComplete = append(contComplete, *c)
			}
			containersservices = contComplete
		}

		common.FormatOutputDef(containersservices)
	},
}
