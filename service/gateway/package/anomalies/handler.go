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
	util.Log(util.Debug, "%s\n", request.URL.RawQuery)

	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if request.Method == http.MethodOptions {
	 	writer.WriteHeader(http.StatusNoContent)
	 	return
	}

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

func Get(values url.Values, key string) []string {
	if values[key] == nil {
		return nil
	}	

	if len(values[key]) == 1 && values[key][0] == "" {
		return nil
	} 

	return values[key]
}

func ParseQuery(values url.Values) *rtr.Filter {
	return &rtr.Filter{
		OpeningDate:           values.Get(OpeningDateKey),
		ClosingDate:           values.Get(ClosingDateKey),
		DistrictName:          Get(values, DistrictNameKey),
		Address:               Get(values, AddressKey),
		ManagementCompanyName: Get(values, ManagementCompanyNameKey),
		UrgencyCategoryName:   Get(values, UrgencyCategoryNameKey),
		AnomalyCategory:       Get(values, AnomalyCategoryKey),
	}
}
