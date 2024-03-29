## hb provision

Provision Healthbot Entities using config files.

### Synopsis

Grouping for a set of commands for provisioning Healthbot Entities.

```
hb provision [flags]
```

### Options

```
  -d, --directory string   Default file location (default "playbook-instances")
  -h, --help               help for provision
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
* [hb provision device-groups](hb_provision_device-groups.md)	 - Provision a set of Device Groups from configuration files.
* [hb provision devices](hb_provision_devices.md)	 - Provision a set of Devices from configuration files.
* [hb provision helper-files](hb_provision_helper-files.md)	 - Upload Helper Files to Healthbot.
* [hb provision playbook-instances](hb_provision_playbook-instances.md)	 - Provision Playbook Instances from configuration files.

###### Auto generated by spf13/cobra on 5-Nov-2019
