package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/dr-livesey-team/service/service/gateway/package/anomalies"
	"github.com/dr-livesey-team/service/service/gateway/package/auth"
	"github.com/dr-livesey-team/service/service/gateway/package/gtw"
	"github.com/dr-livesey-team/service/service/gateway/package/info"
	"github.com/dr-livesey-team/service/service/gateway/package/statistic"
	"github.com/dr-livesey-team/service/service/gateway/package/util"
)

const (
	AnomaliesPath string = "/api/anomalies"
	AuthPath      string = "/api/auth"
	InfoPath      string = "/api/info"
	StatisticPath string = "/api/statistic"
)

func main() {
	config, err := gtw.LoadConfig("config/config.json")
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

	http.Handle(AnomaliesPath, anomalies.NewHandler(config))
	http.Handle(AuthPath, auth.NewHandler(config))
	http.Handle(InfoPath, info.NewHandler(config))
	http.Handle(StatisticPath, statistic.NewHandler(config))

	go http.ListenAndServe(config.Listen, nil)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT)

	<-signals
}
