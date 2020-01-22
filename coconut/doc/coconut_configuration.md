## coconut configuration

view or modify O² configuration

### Synopsis

The configuration command allows you to perform operations on the O² 
configuration store.

### Options

```
  -h, --help   help for configuration
```

### Options inherited from parent commands

```
      --config string            optional configuration file for coconut (default $HOME/.config/coconut/settings.yaml)
      --config_endpoint string   configuration endpoint used by AliECS core as PROTO://HOST:PORT (default "consul://127.0.0.1:8500")
      --endpoint string          AliECS core endpoint as HOST:PORT (default "127.0.0.1:47102")
  -v, --verbose                  show verbose output for debug purposes
```

### SEE ALSO

* [coconut](coconut.md)	 - O² Control and Configuration Utility
* [coconut configuration dump](coconut_configuration_dump.md)	 - dump configuration subtree
* [coconut configuration history](coconut_configuration_history.md)	 - List all existing entries with timestamps of a specified component in Consul
* [coconut configuration import](coconut_configuration_import.md)	 - Import a configuration file for the specified component and entry
* [coconut configuration list](coconut_configuration_list.md)	 - List all existing O² components in Consul
* [coconut configuration show](coconut_configuration_show.md)	 - Show configuration for the component and entry specified

###### Auto generated by spf13/cobra on 20-Jan-2020