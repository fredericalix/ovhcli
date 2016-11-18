package project

import "github.com/spf13/cobra"

func init() {
	Cmd.AddCommand(cmdProjectList)
	Cmd.AddCommand(cmdProjectInfo)

}

// Cmd project
var Cmd = &cobra.Command{
	Use:     "project",
	Short:   "Project commands: ovhcli project --help",
	Long:    `Project commands: ovhcli project <command>`,
	Aliases: []string{"p"},
}
