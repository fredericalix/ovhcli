package connect

import (
	"fmt"
	"os"
	"syscall"

	"github.com/admdwrf/ovhcli/ovhcli/common"
	govh "github.com/ovh/go-ovh/ovh"
	"github.com/spf13/cobra"
)

var endpoint string

func init() {
	Cmd.PersistentFlags().StringVarP(&endpoint, "endpoint", "", "ovh-eu", "ovh-eu | kimsufi-eu | kimsufi-ca | soyoustart-eu | soyoustart-ca | runabove-ca")
}

// Cmd domain
var Cmd = &cobra.Command{
	Use:   "connect",
	Short: "Domain commands: ovhcli connect --help",
	Long:  `Domain commands: ovhcli connect <command>`,
	Run: func(cmd *cobra.Command, args []string) {

		client, err := govh.NewEndpointClient(endpoint)
		if err != nil {
			err = fmt.Errorf("Error while creating OVH Client: %s\nYou need to create an application; please visite this page https://eu.api.ovh.com/createApp/ and create your $HOME/ovh.conf file\n\t[default]\n\t; general configuration: default endpoint\n\tendpoint=ovh-eu\n\n\t[ovh-eu]\n\t; configuration specific to 'ovh-eu' endpoint\n\tapplication_key=my_app_key", err)
		}
		common.Check(err)

		ckReq := client.NewCkRequest()

		// Allow GET method on /me
		ckReq.AddRules(govh.ReadWrite, "/*")

		response, err := ckReq.Do()
		if err != nil {
			fmt.Printf("Error: %q\n", err)
			return
		}

		// set consumer key
		os.Setenv("OVH_CONSUMER_KEY", response.ConsumerKey)

		// Print the validation URL
		fmt.Printf("Please visit %s to complete your login\n", response.ValidationURL)

		syscall.Exec(os.Getenv("SHELL"), []string{os.Getenv("SHELL")}, syscall.Environ())
	},
}
