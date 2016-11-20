package cloud

import (
	"github.com/admdwrf/ovhcli/internal"
	"github.com/spf13/cobra"
)

var cmdProjectList = &cobra.Command{
	Use:   "list",
	Short: "List all projects: ovhcli project list",
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := internal.Client.CloudProjectsList()
		internal.Check(err)
		internal.FormatOutputDef(projects)
	},
}
