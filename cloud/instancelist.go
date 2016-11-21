package cloud

import (
	"github.com/admdwrf/ovhcli/internal"
	"github.com/spf13/cobra"
)

// var withDetails bool

func init() {
	// cmdInstanceList.PersistentFlags().BoolVarP(&withDetails, "withDetails", "", false, "Display cloud instance details")
	cmdInstanceList.PersistentFlags().StringVarP(&projectID, "projectID", "", "", "Your ID Project")
}

var cmdInstanceList = &cobra.Command{
	Use:   "list",
	Short: "List all instance: ovhcli cloud instance list",
	Run: func(cmd *cobra.Command, args []string) {

		instances, err := internal.Client.CloudListInstance(projectID)
		internal.Check(err)

		/*	if withDetails {
			instancesComplete := []sdk.Instance{}
			for _, instance := range instances {
				i, err := internal.Client.VrackInfo(instance.Name)
				internal.Check(err)
				instancesComplete = append(instancesComplete, *i)
			}
			instances = instancesComplete
		} */

		internal.FormatOutputDef(instances)
	},
}
