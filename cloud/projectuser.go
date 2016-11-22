package cloud

import (
	"github.com/spf13/cobra"

	"fmt"

	"strings"

	"github.com/admdwrf/ovhcli/internal"
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
			internal.WrongUsage(cmd)
		},
	}

	cmdProjectUserList = &cobra.Command{
		Use:   "list",
		Short: "List users",
		Run: func(cmd *cobra.Command, args []string) {
			if projectName != "" {
				p, err := internal.Client.CloudProjectInfoByName(projectName)
				internal.Check(err)
				projectID = p.ID
			}

			if projectID == "" {
				internal.WrongUsage(cmd)
			}

			users, err := internal.Client.CloudProjectUsersList(projectID)
			internal.Check(err)
			internal.FormatOutputDef(users)

		},
	}

	cmdProjectCreate = &cobra.Command{
		Use:   "create",
		Short: "Create user",
		Run: func(cmd *cobra.Command, args []string) {
			if projectName != "" {
				p, err := internal.Client.CloudProjectInfoByName(projectName)
				internal.Check(err)
				projectID = p.ID
			}

			if projectID == "" {
				internal.WrongUsage(cmd)
			}

			u, err := internal.Client.CloudProjectUserCreate(projectID, descriptionFlag)
			internal.Check(err)

			regions, err := internal.Client.CloudProjectRegionList(projectID)
			internal.Check(err)

			if envFlag {
				fmt.Println("export OS_AUTH_URL=https://auth.cloud.ovh.net/v2")
				fmt.Printf("# Available regions : %s\n", strings.Join(regions, ", "))
				fmt.Printf("export OS_REGION_NAME=%s\n", regions[0])
				fmt.Printf("export OS_TENANT_ID=%s\n", projectID)
				fmt.Printf("export OS_USERNAME=%s\n", u.Username)
				fmt.Printf("export OS_PASSWORD=%s\n", u.Password)
			}
		},
	}
)
