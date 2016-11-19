package cloud

import (
	"github.com/admdwrf/ovhcli/internal"
	"github.com/spf13/cobra"
)

var projectID string

func init() {
	cmdProjectInfoByID.PersistentFlags().StringVarP(&projectID, "projectID", "", "", "Your ID Project")
}

var cmdProjectInfoByID = &cobra.Command{
	Use:   "infobyid",
	Short: "Info about a project with the project ID: ovhcli cloud project infobyid",
	Run: func(cmd *cobra.Command, args []string) {
		projectinfo, err := internal.Client.CloudProjectInfoByID(projectID)
		internal.Check(err)
		internal.FormatOutputDef(projectinfo)
	},
}
