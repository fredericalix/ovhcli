package vrack

import "github.com/spf13/cobra"

func init() {
	Cmd.AddCommand(cmdVrackList)
}

// Cmd vrack
var Cmd = &cobra.Command{
	Use:   "vrack",
	Short: "Domain commands: ovhcli vrack --help",
	Long:  `Domain commands: ovhcli vrack <command>`,
}
