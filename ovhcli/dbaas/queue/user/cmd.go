package user

import "github.com/spf13/cobra"

var id string
var name string

func init() {
	Cmd.AddCommand(cmdUserList)

	Cmd.PersistentFlags().StringVarP(&id, "id", "", "", "Your Application ID")
	Cmd.PersistentFlags().StringVarP(&name, "name", "", "", "Your Application Name")
}

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "user",
	Short: "Queue user commands: ovhcli dbaas queue user --help",
	Long:  `Queue user commands: ovhcli dbaas queue user <command>`,
}
