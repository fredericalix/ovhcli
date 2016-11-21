
#Configuration 

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

