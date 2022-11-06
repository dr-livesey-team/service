package statistic

import (
	"net/http"
	"net/url"

	"github.com/dr-livesey-team/service/service/gateway/package/anomalies"
	"github.com/dr-livesey-team/service/service/gateway/package/gtw"
	"github.com/dr-livesey-team/service/service/gateway/package/util"
	"github.com/dr-livesey-team/service/service/request_registry/package/rtr"
)

type Handler struct {
	RequestRegistryEndpoint string
}

func NewHandler(config *gtw.Config) *Handler {
	return &Handler{RequestRegistryEndpoint: config.RequestRegistryEndpoint}
}

func (handler *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	filter := ParseQuery(request.URL.Query())

	client, err := rtr.Dial(handler.RequestRegistryEndpoint)
	if err != nil {
		util.LogError(err)
		return
	}

	statistic, err := client.GetStatistic(filter)
	if err != nil {
		util.LogError(err)
		return
	}

	buffer, err := rtr.MarshalStatistic(statistic)
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

func ParseQuery(values url.Values) *rtr.Filter {
	return &rtr.Filter{
		OpeningDate:           values.Get(anomalies.OpeningDateKey),
		ClosingDate:           values.Get(anomalies.ClosingDateKey),
		DistrictName:          values[anomalies.DistrictNameKey],
		Address:               values[anomalies.AddressKey],
		ManagementCompanyName: values[anomalies.ManagementCompanyNameKey],
		UrgencyCategoryName:   values[anomalies.UrgencyCategoryNameKey],
		AnomalyCategory:       values[anomalies.AnomalyCategoryKey],
	}
}
