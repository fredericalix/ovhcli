package cloud

import "github.com/spf13/cobra"

func init() {
	cmdCloudProject.AddCommand(cmdProjectList)
	cmdCloudProject.AddCommand(cmdProjectInfo)
	cmdCloudProject.AddCommand(cmdProjectUser)
}

// cmdCloudProject ...
var cmdCloudProject = &cobra.Command{
	Use:     "project",
	Short:   "Project commands: ovhcli cloud project --help",
	Long:    `Project commands: ovhcli cloud project <command>`,
	Aliases: []string{"pr"},
}
