package cloud

import (
	"github.com/admdwrf/ovhcli/ovhcli/cloud/instance"
	"github.com/admdwrf/ovhcli/ovhcli/cloud/network"
	"github.com/admdwrf/ovhcli/ovhcli/cloud/project"
	"github.com/admdwrf/ovhcli/ovhcli/cloud/sshkey"
	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(project.Cmd)
	Cmd.AddCommand(instance.Cmd)
	Cmd.AddCommand(network.Cmd)
	Cmd.AddCommand(sshkey.Cmd)
}

// Cmd project
var (
	Cmd = &cobra.Command{
		Use:     "cloud",
		Short:   "Project commands: ovhcli cloud --help",
		Long:    `Project commands: ovhcli cloud <command>`,
		Aliases: []string{"cl"},
	}
)
