package topic

import (
	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"

	"github.com/spf13/cobra"
)

var topicID string

func init() {
	cmdInfo.PersistentFlags().StringVarP(&topicID, "topic-id", "", "", "Topic ID")
}

var cmdInfo = &cobra.Command{
	Use:   "info",
	Short: "Get Topic Info: ovhcli dbaas queue topic info (--name=AppName | <--id=appID>) --topic-id=topicid",
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewClient()
		common.Check(err)

		if name != "" {
			app, errInfo := client.DBaasQueueAppInfoByName(name)
			common.Check(errInfo)
			id = app.ID
		}

		topic, err := client.DBaasQueueTopicInfo(id, topicID)
		common.Check(err)

		common.FormatOutputDef(topic)
	},
}
