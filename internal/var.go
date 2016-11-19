package internal

import (
	"github.com/admdwrf/ovhcli/sdk"
)

var (
	// Client is OVHCLI SDK Client
	Client *sdk.Client

	// Format to use for output. One of 'json', 'yaml', 'pretty'
	Format string

	// Verbose ...
	Verbose bool
)
