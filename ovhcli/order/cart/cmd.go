package cart

import "github.com/spf13/cobra"

func init() {
	Cmd.AddCommand(cmdCartList)
	Cmd.AddCommand(cmdCartInfo)
	Cmd.AddCommand(cmdCartAssign)
	Cmd.AddCommand(cmdCartCreate)
	Cmd.AddCommand(cmdCartDelete)
}

// Cmd domain
var Cmd = &cobra.Command{
	Use:   "cart",
	Short: "cart commands: ovhcli order cart --help",
	Long:  `cart commands: ovhcli order cart <command>`,
}
