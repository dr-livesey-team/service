# Overview #

Address registry is an micservice that finds geographical coordinates of buildings. In our particular case it uses
api.mos.data.ru, but in general address registry is easily repacable by other microservice that performs the same work.
It makes our design more flexible.

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
    "listen":"127.0.0.1:50051",
    "pid":"log/main.pid",
    "log":"log/log.txt"
}
```

# Usage #

For example, program that locates address.
```go
package main

import (
    "github.com/dr-livesey-team/service/service/address_registry/package/srv"
    "log"
)

func main() {
    client, err := srv.Dial("127.0.0.0:50050")
    if err != nil {
        log.Fatalln(err.Error())
    }
    defer client.Close()

    request := srv.Request{Address: "Туристская улица, дом 20, корпус 1"}

    response, err := client.Do(&request)
    if err != nil {
        log.Fatalln(err.Error())
    }

    fmt.Prtinf("%g %g\n", response.Latitude, response.Longitude)
}
```