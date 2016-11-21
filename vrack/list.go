package vrack

import (
	"github.com/admdwrf/ovhcli/internal"
	"github.com/admdwrf/ovhcli/sdk"
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

		vracks, err := internal.Client.VrackList()
		internal.Check(err)

		if withDetails {
			vracksComplete := []sdk.Vrack{}
			for _, vrack := range vracks {
				v, err := internal.Client.VrackInfo(vrack.Vrack)
				internal.Check(err)
				vracksComplete = append(vracksComplete, *v)
			}
			vracks = vracksComplete
		}

		internal.FormatOutputDef(vracks)
	},
}
