package queue

import (
	"github.com/admdwrf/ovhcli/ovhcli/dbaas/queue/key"
	"github.com/admdwrf/ovhcli/ovhcli/dbaas/queue/metrics"
	"github.com/admdwrf/ovhcli/ovhcli/dbaas/queue/region"
	"github.com/admdwrf/ovhcli/ovhcli/dbaas/queue/role"
	"github.com/admdwrf/ovhcli/ovhcli/dbaas/queue/service"
	"github.com/admdwrf/ovhcli/ovhcli/dbaas/queue/topic"
	"github.com/admdwrf/ovhcli/ovhcli/dbaas/queue/user"

	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(key.Cmd)
	Cmd.AddCommand(metrics.Cmd)
	Cmd.AddCommand(region.Cmd)
	Cmd.AddCommand(role.Cmd)
	Cmd.AddCommand(service.Cmd)
	Cmd.AddCommand(topic.Cmd)
	Cmd.AddCommand(user.Cmd)
}

// Cmd cmdCloudQueue ...
var Cmd = &cobra.Command{
	Use:   "queue",
	Short: "Queue commands: ovhcli dbaas queue --help",
	Long:  `Queue commands: ovhcli dbaas queue <command>`,
}
