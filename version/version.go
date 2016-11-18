package version

import (
	"fmt"

	"github.com/admdwrf/ovhcli/internal"
	"github.com/spf13/cobra"
)

// Cmd version
var Cmd = &cobra.Command{
	Use:     "version",
	Short:   "Display Version of ovhcli: ovhcli version",
	Long:    `ovhcli version`,
	Aliases: []string{"v"},
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("ovhcli version:", internal.VERSION)

	},
}
