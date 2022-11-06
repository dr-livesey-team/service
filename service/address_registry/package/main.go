package main

import (
	"github.com/dr-livesey-team/service/service/address_registry/package/srv"
	"github.com/dr-livesey-team/service/service/address_registry/package/util"
)

func main() {
	config, err := srv.LoadConfig("config/config.json")
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

	service, err := srv.StartService(config)
	if err != nil {
		util.LogError(err)
	}
	defer service.Stop()

	service.ListenAndServe()
}
