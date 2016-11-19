// ovhcli offers to manage your Ovh services
package main

import (
	"fmt"
	"os"

	"github.com/admdwrf/ovhcli/caas"
	"github.com/admdwrf/ovhcli/cloud"
	"github.com/admdwrf/ovhcli/domain"
	"github.com/admdwrf/ovhcli/internal"
	"github.com/admdwrf/ovhcli/sdk"
	"github.com/admdwrf/ovhcli/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ovhcli",
	Short: "OVH - Command Line Tool",
	Long:  `OVH - Command Line Tool`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error
		internal.Client, err = sdk.NewClient()
		internal.Check(err)
	},
}

func main() {
	rootCmd.PersistentFlags().StringVarP(&internal.Format, "format", "f", "pretty", "choose format output. One of 'json', 'yaml' and 'pretty'")
	rootCmd.PersistentFlags().BoolVarP(&internal.Verbose, "verbose", "v", false, "verbose output")

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

	rootCmd.AddCommand(version.Cmd)
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
