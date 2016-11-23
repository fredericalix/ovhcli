package project

import (
	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

	"github.com/spf13/cobra"
)

func init() {
	cmdProjectRegion.AddCommand(cmdProjectRegionList)
	cmdProjectRegionList.PersistentFlags().BoolVarP(&withDetails, "withDetails", "", false, "Display cloud region details")
}

var (
	cmdProjectRegion = &cobra.Command{
		Use:   "region",
		Short: "Project region management",
		Run: func(cmd *cobra.Command, args []string) {
			common.WrongUsage(cmd)
		},
	}

	cmdProjectRegionList = &cobra.Command{
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
)
