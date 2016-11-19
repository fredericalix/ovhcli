package cloud

import (
	"fmt"
	"os"

	"github.com/admdwrf/ovhcli/sdk"
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

		c, err := sdk.NewClient()
		if err != nil {
			fmt.Printf("Error: %s", err)
			os.Exit(1)
		}

		projectinfo, err := c.CloudProjectInfoByID(projectID)
		if err != nil {
			fmt.Printf("Error: %s", err)
			os.Exit(1)
		}
		fmt.Printf("Project Name: %s\n", projectinfo.Name)
		fmt.Printf("Project Status: %s\n", projectinfo.Status)
		fmt.Printf("Creation Date: %s\n", projectinfo.CreationDate)
	},
}
