package project

import (
	"fmt"
	"os"

	"github.com/admdwrf/ovhcli/sdk/cloud"
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
		// Appel cloud.ProjectInfo(projectid)
		projectinfo, err := cloud.ProjectInfo(projectid)
		if err != nil {
			fmt.Printf("Error: %s", err)
			os.Exit(1)
		}
		fmt.Printf("Project Name: %s\n", projectinfo.Name)
		fmt.Printf("Project Status: %s\n", projectinfo.Status)
		fmt.Printf("Creation Date: %s\n", projectinfo.CreationDate)
	},
}
