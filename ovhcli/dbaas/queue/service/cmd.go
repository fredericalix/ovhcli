package service

import "github.com/spf13/cobra"

var id string
var name string

func init() {
	Cmd.AddCommand(cmdServiceServiceinfo)
	Cmd.AddCommand(cmdServiceList)

	Cmd.PersistentFlags().StringVarP(&id, "id", "", "", "Your Application ID")
	Cmd.PersistentFlags().StringVarP(&name, "name", "", "", "Your Application Name")
}

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "service",
	Short: "Queue service commands: ovhcli dbaas queue service --help",
	Long:  `Queue service commands: ovhcli dbaas queue service <command>`,
}
