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
