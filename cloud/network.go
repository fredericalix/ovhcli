package cloud

import "github.com/spf13/cobra"

func init() {
	cmdCloudNetwork.AddCommand(cmdCloudNetworkPublic)

}

// cmdCloudNetwork ...
var cmdCloudNetwork = &cobra.Command{
	Use:     "network",
	Short:   "Network commands: ovhcli cloud network --help",
	Long:    `Network commands: ovhcli cloud network <command>`,
	Aliases: []string{"net"},
}
