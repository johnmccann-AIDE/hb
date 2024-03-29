## hb scaffold

Generate a config directory from an existing Healthbot installation

### Synopsis

This command when pointed at an existing Healthbot installation, will generate
	valid configuration for the provision sub commands e.g. devices, device-groups, playbook-instances, etc.
	
	The command requires a single argument, the directory where the configs should be written too, current directory is valid.

```
hb scaffold [flags]
```

### Options

```
  -h, --help   help for scaffold
```

### Options inherited from parent commands

```
      --config string     config file (default is $HOME/.hb.yaml)
      --debug             Enable REST debugging
  -p, --password string   Healthbot Password (default "****")
  -r, --resource string   Healthbot Resource Name (default "localhost:8080")
  -u, --username string   Healthbot Username (default "admin")
```

### SEE ALSO

* [hb](hb.md)	 - Healthbot Command Line Interface

###### Auto generated by spf13/cobra on 5-Nov-2019
