package cloud

import (
	"github.com/admdwrf/ovhcli/internal"
	"github.com/spf13/cobra"
)

func init() {
	cmdInstanceInfo.PersistentFlags().StringVarP(&instanceID, "instanceID", "", "", "Your Instance ID")
	cmdInstanceInfo.PersistentFlags().StringVarP(&projectID, "projectID", "", "", "Your ID Project")

}

var cmdInstanceInfo = &cobra.Command{
	Use:   "info",
	Short: "Info about an cloud instance: ovhcli cloud instance info",
	Run: func(cmd *cobra.Command, args []string) {
		instance, err := internal.Client.CloudInfoInstance(projectID, instanceID)
		internal.Check(err)
		internal.FormatOutputDef(instance)
	},
}
