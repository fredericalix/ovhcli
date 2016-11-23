package domain

import "github.com/spf13/cobra"

func init() {
	Cmd.AddCommand(cmdDomainList)
}

// Cmd domain
var Cmd = &cobra.Command{
	Use:   "domain",
	Short: "Domain commands: ovhcli domain --help",
	Long:  `Domain commands: ovhcli domain <command>`,
}
