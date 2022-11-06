# Overview #

We must store requests, anomalies and non anomalies data. For these purpose we chose tarantool database. But, other
microservices doesn't know about it, so that we implement abstraction layer on requests data called requests registry.
Request registry implements necessary functional and easy replaceble by microservice that uses another database. It
makes our design more flexible.

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
will stores logs and pid, database settings such as user, password and endpoint, spaces descriptions and address
registry endpoint. 

For exampl default config.json
```json
{
  "listen":"127.0.0.1:50052",
  "pid":"log/main.pid",
  "log":"log/log.txt",
  "user": "admin",
  "pass": "passwd",
  "database": "127.0.0.1:3301",
  "requests": "spaces/requests.json",
  "anomalies": "spaces/anomalies.json",
  "normal": "spaces/normal.json",
  "adr": "127.0.0.1:50051"
}
```

# Usage #

For example program that selects anomalies that satisfy to given filter.
```go 
package main

import (
    "github.com/dr-livesey-team/service/service/request_registry/package/rtr"
    "log"
)

func main() {
    client, err := rtr.Dial("127.0.0.0:50050")
    if err != nil {
        log.Fatalln(err.Error())
    }
    defer client.Close()

    infos, err := client.GetAnomalies(&rtr.Filter{}) 
    if err != nil {
        log.Fatalln(err.Error())
    }

    fmt.Prtinf("%v\n", infos)
}
```