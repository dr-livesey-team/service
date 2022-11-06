package info

import (
	"math"
	"net/http"
	"net/url"
	"strconv"

	"github.com/dr-livesey-team/service/service/gateway/package/gtw"
	"github.com/dr-livesey-team/service/service/gateway/package/util"
	"github.com/dr-livesey-team/service/service/request_registry/package/rtr"
)

const (
	IdKey string = "id"
)

type Handler struct {
	RequestRegistryEndpoint string
}

func NewHandler(config *gtw.Config) *Handler {
	return &Handler{config.RequestRegistryEndpoint}
}

func (handler *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	util.Log(util.Debug, "%s\n", request.URL.RawQuery)

	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if request.Method == http.MethodOptions {
	 	writer.WriteHeader(http.StatusNoContent)
	 	return
	}

	id := MakeAnomalyIdFromQuery(request.URL.Query())

	client, err := rtr.Dial(handler.RequestRegistryEndpoint)
	if err != nil {
		util.LogError(err)
		return
	}

	infos, err := client.GetRequests(id)
	if err != nil {
		util.LogError(err)
		return
	}

	buffer, err := rtr.MarshalRequestSelectInfos(infos)
	if err != nil {
		util.LogError(err)
		return
	}
	buffer = append(buffer, '\n')

	_, err = writer.Write(buffer)
	if err != nil {
		util.LogError(err)
	}
}

func MakeAnomalyIdFromQuery(values url.Values) *rtr.AnomalyId {
	id, err := strconv.ParseUint(values.Get(IdKey), 10, 64)
	if err != nil {
		util.LogError(err)
		return &rtr.AnomalyId{
			Id: math.MaxUint64,
		}
	}

	return &rtr.AnomalyId{
		Id: id,
	}
}
