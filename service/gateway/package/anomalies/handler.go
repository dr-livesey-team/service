package anomalies

import (
	"net/http"
	"net/url"

	"github.com/dr-livesey-team/service/service/gateway/package/gtw"
	"github.com/dr-livesey-team/service/service/gateway/package/util"
	"github.com/dr-livesey-team/service/service/request_registry/package/rtr"
)

const (
	OpeningDateKey           string = "opening_date"
	ClosingDateKey           string = "closing_date"
	DistrictNameKey          string = "district_name"
	AddressKey               string = "address"
	ManagementCompanyNameKey string = "management_company_name"
	UrgencyCategoryNameKey   string = "urgency_category_name"
	AnomalyCategoryKey       string = "anomaly_category"
)

type Handler struct {
	RequestRegistryEndpoint string
}

func NewHandler(config *gtw.Config) *Handler {
	return &Handler{config.RequestRegistryEndpoint}
}

func (handler *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	filter := ParseQuery(request.URL.Query())

	client, err := rtr.Dial(handler.RequestRegistryEndpoint)
	if err != nil {
		util.LogError(err)
		return
	}

	infos, err := client.GetAnomalies(filter)
	if err != nil {
		util.LogError(err)
		return
	}

	buffer, err := rtr.MarshalAnomalySelectInfos(infos)
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
		OpeningDate:           values.Get(OpeningDateKey),
		ClosingDate:           values.Get(ClosingDateKey),
		DistrictName:          values[DistrictNameKey],
		Address:               values[AddressKey],
		ManagementCompanyName: values[ManagementCompanyNameKey],
		UrgencyCategoryName:   values[UrgencyCategoryNameKey],
		AnomalyCategory:       values[AnomalyCategoryKey],
	}
}
