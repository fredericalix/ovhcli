package cloud

import (
	"github.com/admdwrf/ovhcli/internal"
	"github.com/admdwrf/ovhcli/sdk"
	"github.com/spf13/cobra"
)

var withDetails bool

func init() {
	cmdProjectList.PersistentFlags().BoolVarP(&withDetails, "withDetails", "", false, "Display project details")
}

var cmdProjectList = &cobra.Command{
	Use:   "list",
	Short: "List all projects: ovhcli project list",
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := internal.Client.CloudProjectsList()

		if withDetails {
			projectsComplete := []sdk.Project{}
			for _, project := range projects {
				p, e := internal.Client.CloudProjectInfoByID(project.ID)
				internal.Check(e)
				projectsComplete = append(projectsComplete, *p)
			}
			projects = projectsComplete
		}

		internal.Check(err)
		internal.FormatOutputDef(projects)
	},
}
