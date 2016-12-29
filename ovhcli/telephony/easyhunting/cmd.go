package easyhunting

import (
	"github.com/admdwrf/ovhcli/ovhcli/telephony/easyhunting/hunting"

	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(cmdEasyHuntingList)
	Cmd.AddCommand(cmdEasyHuntingInfo)

	Cmd.AddCommand(hunting.Cmd)
}

// Cmd ...
var Cmd = &cobra.Command{
	Use:   "easyhunting",
	Short: "EasyHunting commands: ovhcli telephony easyhunting --help",
	Long:  `EasyHunting commands: ovhcli telephony easyhunting <command>`,
}
