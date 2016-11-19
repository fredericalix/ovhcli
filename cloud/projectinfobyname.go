package cloud

import (
	"fmt"
	"os"

	"github.com/admdwrf/ovhcli/sdk"
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

		c, err := sdk.NewClient()
		if err != nil {
			fmt.Printf("Error: %s", err)
			os.Exit(1)
		}

		projectinfo, err := c.CloudProjectInfoByName(projectName)
		if err != nil {
			fmt.Printf("Error: %s", err)
			os.Exit(1)
		}
		fmt.Printf("Project Name: %s\n", projectinfo.Name)
		fmt.Printf("Project Status: %s\n", projectinfo.Status)
		fmt.Printf("Creation Date: %s\n", projectinfo.CreationDate)
	},
}
