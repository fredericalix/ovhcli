package cloud

import (
	"fmt"
	"strings"

	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

	"github.com/spf13/cobra"
)

func init() {
	cmdProjectUser.PersistentFlags().StringVarP(&projectID, "id", "", "", "Your ID Project")
	cmdProjectUser.PersistentFlags().StringVarP(&projectName, "name", "", "", "Your Project Name")
	cmdProjectUser.AddCommand(cmdProjectUserList)
	cmdProjectUser.AddCommand(cmdProjectCreate)

	cmdProjectCreate.Flags().BoolVarP(&envFlag, "env", "", false, "Helps to eval printed values as standard OpenStack environment variables")
	cmdProjectCreate.Flags().StringVarP(&descriptionFlag, "description", "", "", "User description")
}

var (
	envFlag         bool
	descriptionFlag string

	cmdProjectUser = &cobra.Command{
		Use:   "user",
		Short: "Project user management",
		Run: func(cmd *cobra.Command, args []string) {
			common.WrongUsage(cmd)
		},
	}

	cmdProjectUserList = &cobra.Command{
		Use:   "list",
		Short: "List users",
		Run: func(cmd *cobra.Command, args []string) {
			client, err := ovh.NewClient()
			common.Check(err)

			if projectName != "" {
				p, err := client.CloudProjectInfoByName(projectName)
				common.Check(err)
				projectID = p.ID
			}

			if projectID == "" {
				common.WrongUsage(cmd)
			}

			users, err := client.CloudProjectUsersList(projectID)
			common.Check(err)
			common.FormatOutputDef(users)

		},
	}

	cmdProjectCreate = &cobra.Command{
		Use:   "create",
		Short: "Create user",
		Run: func(cmd *cobra.Command, args []string) {
			client, err := ovh.NewClient()
			common.Check(err)

			if projectName != "" {
				p, err := client.CloudProjectInfoByName(projectName)
				common.Check(err)
				projectID = p.ID
			}

			if projectID == "" {
				common.WrongUsage(cmd)
			}

			u, err := client.CloudProjectUserCreate(projectID, descriptionFlag)
			common.Check(err)

			if envFlag {
				regions, err := client.CloudProjectRegionList(projectID)
				common.Check(err)

				fmt.Println("export OS_AUTH_URL=https://auth.cloud.ovh.net/v2")
				fmt.Printf("# Available regions : %s\n", strings.Join(regions, ", "))
				fmt.Printf("export OS_REGION_NAME=%s\n", regions[0])
				fmt.Printf("export OS_TENANT_ID=%s\n", projectID)
				fmt.Printf("export OS_USERNAME=%s\n", u.Username)
				fmt.Printf("export OS_PASSWORD=%s\n", u.Password)
				return
			}

			common.FormatOutputDef(u)
		},
	}
)
