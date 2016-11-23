package vrack

import (
	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

	"github.com/spf13/cobra"
)

var withDetails bool

func init() {
	cmdVrackList.PersistentFlags().BoolVarP(&withDetails, "withDetails", "", false, "Display vrack details")
}

var cmdVrackList = &cobra.Command{
	Use:   "list",
	Short: "List all vrack: ovhcli vrack list",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewClient()
		common.Check(err)

		vracks, err := client.VrackList()
		common.Check(err)

		if withDetails {
			vracksComplete := []ovh.Vrack{}
			for _, vrack := range vracks {
				v, err := client.VrackInfo(vrack.Name)
				common.Check(err)
				vracksComplete = append(vracksComplete, *v)
			}
			vracks = vracksComplete
		}

		common.FormatOutputDef(vracks)
	},
}
