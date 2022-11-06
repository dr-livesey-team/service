# Overview #

API for front end and another applications that interested in requsts data.

# Build and launch #

To build address registry run
```
$ make init
$ make
$ make utils
```

To clean executables, utilitiess and auxiliary files run
```
$ make clean
```

To launch address registry
```
$ launcher/launcher main
```

# Configuration #

Configuration file uses to set up address registry work. You can specify endpoint which servis will listen, file where
will stores logs and pid.

For exampl default config.json
```json
{
  "listen":"127.0.0.1:50050",
  "pid":"log/main.pid",
  "log":"log/log.txt",
  "rtr": "127.0.0.1:50052"
}
```
