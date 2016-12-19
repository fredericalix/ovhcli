package project

import (
	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

	"github.com/spf13/cobra"
)


func init() {
	cmdProjectList.Flags().BoolVarP(&withDetails, "withDetails", "", false, "Display project details")
}

var cmdProjectList = &cobra.Command{
	Use:   "list",
	Short: "List all projects: ovhcli project list",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := ovh.NewClient()
		common.Check(err)

		projects, err := client.CloudProjectsList()

		if withDetails {
			projectsComplete := []ovh.Project{}
			for _, project := range projects {
				p, e := client.CloudProjectInfoByID(project.ID)
				common.Check(e)
				projectsComplete = append(projectsComplete, *p)
			}
			projects = projectsComplete
		}

		common.Check(err)
		common.FormatOutputDef(projects)
	},
}
