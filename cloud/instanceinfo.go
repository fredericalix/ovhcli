package cloud

import (
	"github.com/admdwrf/ovhcli/internal"
	"github.com/spf13/cobra"
)

var instanceName string

func init() {
	cmdInstanceInfo.PersistentFlags().StringVarP(&instanceName, "instanceName", "", "", "Your Instance name")
	cmdInstanceInfo.PersistentFlags().StringVarP(&projectID, "projectID", "", "", "Your ID Project")

}

var cmdInstanceInfo = &cobra.Command{
	Use:   "info",
	Short: "Info about an cloud instance: ovhcli cloud instance info",
	Run: func(cmd *cobra.Command, args []string) {
		instance, err := internal.Client.CloudInfoInstance(projectID, instanceName)
		internal.Check(err)
		internal.FormatOutputDef(instance)
	},
}
