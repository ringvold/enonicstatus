Enonicstatus
============

Enonicstatus is a commandline application that displays informastion about
an Enonic cluster.

Currently the tool displayes a fixed set of fields (name, index status,
master status, nodes seen, uptime and enonic version) but there are plans
to make this customizable.


## Usage

`enonicstatus cms --hosts=[list of hosts]`

For more details see `enonicstatus -h`

## Configure

In addition to the flags the command looks for a file called .enonicstatus.yaml
(TOML, JAON and HCL also supported thanks to [Viper](https://github.com/spf13/viper)) in the current directory and home directory
for configuration. This makes it easier to use and switch between multiple
environments (clusters).

.enonicstatus.yaml example:
``` YAML
noProxy: true
prod:
  hosts: "serverprod1, serverprod2, serverprod3, serverprod4"
qa:
  hosts: "serverqa1, serverqa2, serverqa3, serverqa4"
```

Note: When using cygwin the program has problems finding the config file in
your home directory so it needs to be specified, here with a "mixed" path:
`enonicstatus cms --config="\Users\myuser\.enonicstatus.yaml"`. It does find the
file in the current directory.
