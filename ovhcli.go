// ovhcli offers to manage your Ovh services
package main

import (
	"fmt"
	"os"

	"github.com/admdwrf/ovhcli/project"
	"github.com/admdwrf/ovhcli/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ovhcli",
	Short: "OVH - Command Line Tool",
	Long:  `OVH - Command Line Tool`,
}

func main() {
	addCommands()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

//AddCommands adds child commands to the root command rootCmd.
func addCommands() {
	rootCmd.AddCommand(project.Cmd)
	rootCmd.AddCommand(version.Cmd)
}
