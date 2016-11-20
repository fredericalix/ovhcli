package cloud

import "github.com/spf13/cobra"

func init() {
	cmdCloudInstance.AddCommand(cmdInstanceDelete)
	cmdCloudInstance.AddCommand(cmdInstanceCreate)

}

// cmdCloudInstance ...
var cmdCloudInstance = &cobra.Command{
	Use:     "instance",
	Short:   "Instance commands: ovhcli cloud instance --help",
	Long:    `Instance commands: ovhcli cloud instance <command>`,
	Aliases: []string{"in"},
}
