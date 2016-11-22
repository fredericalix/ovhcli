package cloud

import "github.com/spf13/cobra"

func init() {
	cmdCloudRegion.AddCommand(cmdCloudRegionList)

}

// cmdCloudInstance ...
var cmdCloudRegion = &cobra.Command{
	Use:     "region",
	Short:   "Region commands: ovhcli cloud region --help",
	Long:    `Region commands: ovhcli cloud region <command>`,
	Aliases: []string{"re"},
}
