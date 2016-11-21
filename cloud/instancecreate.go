package cloud

import (
	"github.com/admdwrf/ovhcli/internal"
	"github.com/spf13/cobra"
)

var projectID string
var imageID string
var name string
var pubkeyID string
var flavorID string
var region string

func init() {
	cmdInstanceCreate.PersistentFlags().StringVarP(&projectID, "projectID", "", "", "Your ID Project")
	cmdInstanceCreate.PersistentFlags().StringVarP(&name, "name", "", "", "Your Instance name to create")
	cmdInstanceCreate.PersistentFlags().StringVarP(&imageID, "imageID", "", "", "Your image ID to use")
	cmdInstanceCreate.PersistentFlags().StringVarP(&pubkeyID, "pubkeyID", "", "", "Your sshkey ID to use")
	cmdInstanceCreate.PersistentFlags().StringVarP(&flavorID, "flavorID", "", "", "Your flavor ID to use")
	cmdInstanceCreate.PersistentFlags().StringVarP(&region, "region", "", "", "region to use")

}

var cmdInstanceCreate = &cobra.Command{
	Use:   "create",
	Short: "Create Cloud Public Instance: ovhcli cloud instance create",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := internal.Client.CloudCreateInstance(projectID, name, pubkeyID, flavorID, imageID, region)
		internal.Check(err)
		internal.FormatOutputDef(c)
	},
}
