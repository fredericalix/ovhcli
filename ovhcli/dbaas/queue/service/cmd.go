package service

import "github.com/spf13/cobra"

func init() {
	Cmd.AddCommand(cmdServiceList)
}

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "service",
	Short: "Queue service commands: ovhcli dbaas queue service --help",
	Long:  `Queue service commands: ovhcli dbaas queue service <command>`,
}
