package cloud

import (
	"github.com/admdwrf/ovhcli/internal"
	"github.com/admdwrf/ovhcli/sdk"
	"github.com/spf13/cobra"
)

// var withDetails bool

func init() {
	cmdCloudRegionList.PersistentFlags().BoolVarP(&withDetails, "withDetails", "", false, "Display cloud region details")
	cmdCloudRegionList.PersistentFlags().StringVarP(&projectID, "projectID", "", "", "Your ID Project")
}

var cmdCloudRegionList = &cobra.Command{
	Use:   "list",
	Short: "List all regions: ovhcli cloud region list",
	Run: func(cmd *cobra.Command, args []string) {

		regions, err := internal.Client.CloudListRegions(projectID)
		internal.Check(err)

		if withDetails {
			regionsComplete := []sdk.Regions{}
			for _, region := range regions {
				r, err := internal.Client.CloudInfoRegion(projectID, region.Region)
				internal.Check(err)
				regionsComplete = append(regionsComplete, *r)
			}
			regions = regionsComplete

		}

		internal.FormatOutputDef(regions)
	},
}
