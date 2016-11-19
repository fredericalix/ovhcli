package project

import (
	"github.com/admdwrf/ovhcli/internal"
	"github.com/spf13/cobra"
)

var projectid string

func init() {
	cmdProjectInfo.PersistentFlags().StringVarP(&projectid, "projectid", "", "", "Your ID Project")
}

var cmdProjectInfo = &cobra.Command{
	Use:   "info",
	Short: "Info about a project: ovhcli project info",
	Run: func(cmd *cobra.Command, args []string) {
		projectinfo, err := internal.Client.CloudProjectInfo(projectid)
		internal.Check(err)
		internal.FormatOutputDef(projectinfo)
	},
}
