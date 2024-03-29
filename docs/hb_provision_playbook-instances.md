## hb provision playbook-instances

Provision Playbook Instances from configuration files.

### Synopsis

The Playbook Instances can be defined in YAML or JSON and conform to the payload definitions for the REST API.

```
hb provision playbook-instances [flags]
```

### Options

```
  -h, --help   help for playbook-instances
```

### Options inherited from parent commands

```
      --config string      config file (default is $HOME/.hb.yaml)
      --debug              Enable REST debugging
  -d, --directory string   Default file location (default "playbook-instances")
  -p, --password string    Healthbot Password (default "****")
  -r, --resource string    Healthbot Resource Name (default "localhost:8080")
  -u, --username string    Healthbot Username (default "admin")
```

### SEE ALSO

* [hb provision](hb_provision.md)	 - Provision Healthbot Entities using config files.

###### Auto generated by spf13/cobra on 5-Nov-2019
