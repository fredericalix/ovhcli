package cloud

import (
	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

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
		client, err := ovh.NewClient()
		common.Check(err)

		regions, err := client.CloudListRegions(projectID)
		common.Check(err)

		if withDetails {
			regionsComplete := []ovh.Regions{}
			for _, region := range regions {
				r, err := client.CloudInfoRegion(projectID, region.Region)
				common.Check(err)
				regionsComplete = append(regionsComplete, *r)
			}
			regions = regionsComplete

		}

		common.FormatOutputDef(regions)
	},
}
