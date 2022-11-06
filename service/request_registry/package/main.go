package main

import (
	"github.com/elisfromkirov/service/service/request_registry/package/rtr"
	"github.com/elisfromkirov/service/service/request_registry/package/util"
)

func main() {
	config, err := rtr.LoadConfig("config/config.json")
	if err != nil {
		panic(err.Error())
	}

	err = util.OpenLogFile(config.LogFile)
	if err != nil {
		panic(err.Error())
	}
	defer util.CloseLogFile()

	err = util.LogPid(config.PidFile)
	if err != nil {
		panic(err.Error())
	}

	service, err := rtr.StartRequestRegistry(config)
	if err != nil {
		panic(err.Error())
	}
	defer service.Stop()

	server, err := rtr.StartServer(config.Listen, service)
	if err != nil {
		panic(err.Error())
	}
	defer server.Stop()

	server.ListenAndServe()
}
