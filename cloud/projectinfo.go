package cloud

import (
	"github.com/spf13/cobra"

	"github.com/admdwrf/ovhcli/internal"
	"github.com/admdwrf/ovhcli/sdk"
)

func init() {
	cmdProjectInfo.PersistentFlags().StringVarP(&projectID, "id", "", "", "Your ID Project")
	cmdProjectInfo.PersistentFlags().StringVarP(&projectName, "name", "", "", "Your Project Name")
}

var (
	projectName string

	cmdProjectInfo = &cobra.Command{
		Use:   "info",
		Short: "Info about a project",
		Run: func(cmd *cobra.Command, args []string) {
			if projectID == "" && projectName == "" {
				internal.WrongUsage(cmd)
			}
			var p *sdk.Project
			var err error
			if projectID != "" {
				p, err = internal.Client.CloudProjectInfoByID(projectID)
			} else {
				p, err = internal.Client.CloudProjectInfoByName(projectName)
			}
			internal.Check(err)
			internal.FormatOutputDef(p)
		},
	}
)
