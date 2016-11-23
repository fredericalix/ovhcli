package project

import "github.com/spf13/cobra"

func init() {
	Cmd.AddCommand(cmdProjectList)
	Cmd.AddCommand(cmdProjectInfo)
	Cmd.AddCommand(cmdProjectUser)
	Cmd.AddCommand(cmdProjectRegion)

	Cmd.PersistentFlags().StringVarP(&projectID, "id", "", "", "Your ID Project")
	Cmd.PersistentFlags().StringVarP(&projectName, "name", "", "", "Your Project Name")
}

// Cmd ...
var (
	projectID   string
	projectName string
	withDetails bool

	Cmd = &cobra.Command{
		Use:     "project",
		Short:   "Project commands: ovhcli cloud project --help",
		Long:    `Project commands: ovhcli cloud project <command>`,
		Aliases: []string{"pr"},
	}
)
