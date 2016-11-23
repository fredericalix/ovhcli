package cloud

import (
	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

	"github.com/spf13/cobra"
)

var project string

func init() {
	cmdCloudNetworkPublicShow.PersistentFlags().StringVarP(&project, "project", "", "", "Your ID Project")
}

// cmdCloudNetworkPublicShow show Public network ID of a project
var cmdCloudNetworkPublicShow = &cobra.Command{
	Use:   "show",
	Short: "Show the public network ID of your project: ovhcli cloud network public show",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ovh.NewClient()
		common.Check(err)

		netpub, err := client.CloudInfoNetworkPublic(project)

		common.Check(err)
		common.FormatOutputDef(netpub)
	},
}
