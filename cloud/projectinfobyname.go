package cloud

import (
	"github.com/admdwrf/ovhcli/internal"
	"github.com/spf13/cobra"
)

var projectName string

func init() {
	cmdProjectInfoByName.PersistentFlags().StringVarP(&projectName, "projectname", "", "", "Your Project name")
}

var cmdProjectInfoByName = &cobra.Command{
	Use:   "infobyname",
	Short: "Info about a project with the project name: ovhcli cloud project infobyname",
	Run: func(cmd *cobra.Command, args []string) {
		projectinfo, err := internal.Client.CloudProjectInfoByName(projectName)
		internal.Check(err)
		internal.FormatOutputDef(projectinfo)
	},
}
