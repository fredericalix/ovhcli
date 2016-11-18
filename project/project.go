package project

import "github.com/spf13/cobra"

func init() {
	Cmd.AddCommand(cmdProjectList)
}

// Cmd project
var Cmd = &cobra.Command{
	Use:     "project",
	Short:   "Project commands: tatcli project --help",
	Long:    `Project commands: tatcli project <command>`,
	Aliases: []string{"u"},
}
