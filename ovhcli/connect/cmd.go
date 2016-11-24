package connect

import (
	"fmt"
	"os"
	"syscall"

	ovh "github.com/admdwrf/ovhcli"
	"github.com/admdwrf/ovhcli/ovhcli/common"
	govh "github.com/ovh/go-ovh/ovh"
	"github.com/spf13/cobra"
)

// Cmd domain
var Cmd = &cobra.Command{
	Use:   "connect",
	Short: "Domain commands: ovhcli connect --help",
	Long:  `Domain commands: ovhcli connect <command>`,
	Run: func(cmd *cobra.Command, args []string) {

		client, err := ovh.NewClient()
		common.Check(err)

		ckReq := client.OVHClient.NewCkRequest()

		// Allow GET method on /me
		ckReq.AddRules(govh.ReadWrite, "/*")

		response, err := ckReq.Do()
		common.Check(err)

		// set consumer key
		os.Setenv("OVH_CONSUMER_KEY", response.ConsumerKey)

		// Print the validation URL
		fmt.Printf("Please visit %s to complete your login\n", response.ValidationURL)

		syscall.Exec(os.Getenv("SHELL"), []string{os.Getenv("SHELL")}, syscall.Environ())
	},
}
