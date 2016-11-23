package cloud

import "github.com/spf13/cobra"

func init() {
	cmdCloudSSHkey.AddCommand(cmdCloudSSHKeyList)

}

// cmdCloudSSHkey ...
var cmdCloudSSHkey = &cobra.Command{
	Use:     "sshkey",
	Short:   "sshkey commands: ovhcli cloud sshkey --help",
	Long:    `Regisshkeyon commands: ovhcli cloud sshkey <command>`,
	Aliases: []string{"ssh"},
}
