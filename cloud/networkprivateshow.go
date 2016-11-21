package cloud

import (
	"github.com/admdwrf/ovhcli/internal"
	"github.com/spf13/cobra"
)

//var project string

func init() {
	cmdCloudNetworkPrivateShow.PersistentFlags().StringVarP(&project, "project", "", "", "Your ID Project")
}

// cmdCloudNetworkPrivateShow show Public network ID of a project
var cmdCloudNetworkPrivateShow = &cobra.Command{
	Use:   "show",
	Short: "Show the private network ID of your project: ovhcli cloud network private show",
	Run: func(cmd *cobra.Command, args []string) {
		netpub, err := internal.Client.CloudInfoNetworkPrivate(project)

		internal.Check(err)
		internal.FormatOutputDef(netpub)
	},
}
