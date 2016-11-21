package cloud

import (
	"github.com/admdwrf/ovhcli/internal"
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
		netpub, err := internal.Client.CloudInfoNetworkPublic(project)

		internal.Check(err)
		internal.FormatOutputDef(netpub)
	},
}
