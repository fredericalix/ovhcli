// ovhcli offers to manage your Ovh services
package main

import (
	"fmt"
	"os"

	"github.com/admdwrf/ovhcli/ovhcli/caas"
	"github.com/admdwrf/ovhcli/ovhcli/cloud"
	"github.com/admdwrf/ovhcli/ovhcli/common"
	"github.com/admdwrf/ovhcli/ovhcli/connect"
	"github.com/admdwrf/ovhcli/ovhcli/dbaas"
	"github.com/admdwrf/ovhcli/ovhcli/domain"
	"github.com/admdwrf/ovhcli/ovhcli/version"
	"github.com/admdwrf/ovhcli/ovhcli/vrack"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ovhcli",
	Short: "OVH - Command Line Tool",
}

func main() {
	rootCmd.PersistentFlags().StringVarP(&common.Format, "format", "f", "pretty", "choose format output. One of 'json', 'yaml' and 'pretty'")
	rootCmd.PersistentFlags().BoolVarP(&common.Verbose, "verbose", "v", false, "verbose output")

	addCommands()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

//AddCommands adds child commands to the root command rootCmd.
func addCommands() {
	rootCmd.AddCommand(caas.Cmd)
	rootCmd.AddCommand(domain.Cmd)
	rootCmd.AddCommand(cloud.Cmd)
	rootCmd.AddCommand(dbaas.Cmd)

	rootCmd.AddCommand(version.Cmd)
	rootCmd.AddCommand(vrack.Cmd)

	rootCmd.AddCommand(connect.Cmd)

	rootCmd.AddCommand(autocompleteCmd)
}

var autocompleteCmd = &cobra.Command{
	Use:   "autocomplete <path>",
	Short: "Generate bash autocompletion file for ovhcli",
	Long:  `Generate bash autocompletion file for ovhcli`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Fprintf(os.Stderr, "Wrong usage: ovhcli autocomplete <path>\n")
			return
		}
		rootCmd.GenBashCompletionFile(args[0])
		fmt.Fprintf(os.Stderr, "Completion file generated.\n")
		fmt.Fprintf(os.Stderr, "You may now run `source %s`\n", args[0])
	},
}
