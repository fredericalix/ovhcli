package project

import (
	"github.com/spf13/cobra"

	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"
)

var (
	cmdProjectInfo = &cobra.Command{
		Use:   "info",
		Short: "Info about a project",
		Run: func(cmd *cobra.Command, args []string) {
			client, err := ovh.NewClient()
			common.Check(err)

			if projectID == "" && projectName == "" {
				common.WrongUsage(cmd)
			}
			var p *ovh.Project
			if projectID != "" {
				p, err = client.CloudProjectInfoByID(projectID)
			} else {
				p, err = client.CloudProjectInfoByName(projectName)
			}
			common.Check(err)
			common.FormatOutputDef(p)
		},
	}
)
