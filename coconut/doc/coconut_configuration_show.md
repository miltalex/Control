## coconut configuration show

Show configuration for the component and entry specified

### Synopsis

The configuration show command returns the most recent 
configuration revision for the specified component and entry. 
It can also return a specific revision, requested with the --timestamp/-t flag

```
coconut configuration show <component> <entry> [flags]
```

### Examples

```
coconut conf show <component> <entry> 
coconut conf show <component> <entry> -t <timestamp>
coconut conf show <component>/<entry>
coconut conf show <component>/<entry> -t <timestamp>
coconut conf show <component>/<entry>@<timestamp>
```

### Options

```
  -h, --help               help for show
  -t, --timestamp string   request configuration for this timestamp
```

### Options inherited from parent commands

```
      --config string            optional configuration file for coconut (default $HOME/.config/coconut/settings.yaml)
      --config_endpoint string   configuration endpoint used by AliECS core as PROTO://HOST:PORT (default "consul://127.0.0.1:8500")
      --endpoint string          AliECS core endpoint as HOST:PORT (default "127.0.0.1:47102")
  -v, --verbose                  show verbose output for debug purposes
```

### SEE ALSO

* [coconut configuration](coconut_configuration.md)	 - view or modify O² configuration

###### Auto generated by spf13/cobra on 20-Jan-2020