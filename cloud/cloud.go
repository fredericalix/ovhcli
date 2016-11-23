package cloud

import "github.com/spf13/cobra"

func init() {
	Cmd.AddCommand(cmdCloudProject)
	Cmd.AddCommand(cmdCloudInstance)
	Cmd.AddCommand(cmdCloudRegion)
	Cmd.AddCommand(cmdCloudNetwork)
	Cmd.AddCommand(cmdCloudSSHkey)

}

// Cmd project
var Cmd = &cobra.Command{
	Use:     "cloud",
	Short:   "Project commands: ovhcli cloud --help",
	Long:    `Project commands: ovhcli cloud <command>`,
	Aliases: []string{"cl"},
}
