[![GoDoc](https://godoc.org/github.com/admdwrf/ovhcli?status.svg)](https://godoc.org/github.com/admdwrf/ovhcli)
[![Go Report Card](https://goreportcard.com/badge/admdwrf/ovhcli)](https://goreportcard.com/report/admdwrf/ovhcli)

# Build

```bash
mkdir -p $GOPATH/github.com/admdwrf/
cd $GOPATH/github.com/admdwrf/
git clone git@github.com:admdwrf/ovhcli.git
cd $GOPATH/github.com/admdwrf/ovhcli/ovhcli
go build
./ovhcli -h
```

# Configuration

ovhcli uses [go-ovh](https://github.com/ovh/go-ovh) to connect on api.

Before run cli, you need set environment variables :

- ``OVH_ENDPOINT``,
- ``OVH_APPLICATION_KEY``,
- ``OVH_APPLICATION_SECRET``
- ``OVH_CONSUMER_KEY``  

If either of these parameter is not provided, it will look for a configuration file
at these paths :

- ./ovh.conf
- $HOME/.ovh.conf
- /etc/ovh.conf

```ini
[default]
; general configuration: default endpoint
endpoint=ovh-eu

[ovh-eu]
; configuration specific to 'ovh-eu' endpoint
application_key=my_app_key
application_secret=my_application_secret
consumer_key=my_consumer_key
```

For more information about configuration : https://github.com/ovh/go-ovh

# Use SDK

Example for listing domains

```go
package main

import (
	"fmt"
	ovh "github.com/admdwrf/ovhcli"
)

func main() {
	client, err := ovh.NewClient()
	if err != nil {
		fmt.Printf("Error:%s", err)
	}

	domains, err := client.DomainList()
	if err != nil {
		fmt.Printf("Error:%s", err)
	}

	for _, domain := range domains {
		fmt.Printf("Domain:%s", domain.Domain)
	}
}
```
