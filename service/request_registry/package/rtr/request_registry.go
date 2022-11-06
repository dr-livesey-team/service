package rtr

import (
	"github.com/dr-livesey-team/service/service/address_registry/package/srv"
	"github.com/dr-livesey-team/service/service/request_registry/package/date"
	"github.com/dr-livesey-team/service/service/request_registry/package/util"
	"github.com/tarantool/go-tarantool"
)

const (
	SelectAnomalies string = "select_anomalies"
	NumberAnomalies string = "number_anomalies"
	SelectRequests  string = "select_requests"
	NumberNormal    string = "number_normal"
)

const (
	OpeningDateKey           string = "opening_date"
	ClosingDateKey           string = "closing_date"
	DistrictNameKey          string = "district_name"
	AddressKey               string = "address"
	ManagementCompanyNameKey string = "management_company_name"
	ServiceOrganizationKey   string = "service_organization_name"
	UrgencyCategoryNameKey   string = "urgency_category_name"
)

const (
	MinDate string = "2021-01-01"
	MaxDate string = "2021-12-31"
)

type RequestRegistry struct {
	Conn            *tarantool.Connection
	Requests        *SpaceDesc
	Anomalies       *SpaceDesc
	Normal          *SpaceDesc
	AddressRegistry string
}

func StartRequestRegistry(config *Config) (*RequestRegistry, error) {
	requests, err := LoadSpaceDesc(config.RequestsDescFile)
	if err != nil {
		return nil, err
	}

	anomalies, err := LoadSpaceDesc(config.AnomaliesDescFile)
	if err != nil {
		return nil, err
	}

	normal, err := LoadSpaceDesc(config.NormalDescFile)
	if err != nil {
		return nil, err
	}

	conn, err := tarantool.Connect(config.DataBase, tarantool.Opts{
		User: config.User,
		Pass: config.Pass,
	})
	if err != nil {
		return nil, err
	}

	return &RequestRegistry{conn, requests, anomalies, normal, config.AddressRegistry}, nil
}

func (service *RequestRegistry) Stop() error {
	return service.Conn.Close()
}

func (service *RequestRegistry) GetAnomalies(filter *Filter) (*AnomalySelectInfos, error) {
	if filter.OpeningDate == "" {
		filter.OpeningDate = MinDate
	}

	if filter.ClosingDate == "" {
		filter.ClosingDate = MaxDate
	}

	var infos []AnomalySelectInfo

	err := service.Conn.CallTyped(SelectAnomalies,
		[]interface{}{
			map[string]interface{}{
				OpeningDateKey:           filter.OpeningDate,
				ClosingDateKey:           filter.ClosingDate,
				DistrictNameKey:          filter.DistrictName,
				AddressKey:               filter.Address,
				ManagementCompanyNameKey: filter.ManagementCompanyName,
				ServiceOrganizationKey:   filter.ServiceOrganizationName,
				UrgencyCategoryNameKey:   filter.UrgencyCategoryName,
			},
		},
		&infos,
	)
	if err != nil {
		util.LogError(err)
		return nil, err
	}

	return &AnomalySelectInfos{Infos: infos}, nil
}

func (service *RequestRegistry) GetRequests(info *AnomalyId) (*RequestSelectInfos, error) {
	var infos []RequestSelectInfo

	err := service.Conn.CallTyped(SelectRequests,
		[]interface{}{info.Id},
		&infos,
	)
	if err != nil {
		util.LogError(err)
		return nil, err
	}

	return &RequestSelectInfos{Infos: infos}, nil
}

type NumberTuple struct {
	Value uint32
}

func (service *RequestRegistry) GetStatistic(filter *Filter) (*Statistic, error) {
	if filter.OpeningDate == "" {
		filter.OpeningDate = MinDate
	}

	if filter.ClosingDate == "" {
		filter.ClosingDate = MaxDate
	}

	var points []Point

	for filter.OpeningDate != filter.ClosingDate {
		var numberAnomalies []NumberTuple

		err := service.Conn.CallTyped(NumberAnomalies,
			[]interface{}{
				map[string]interface{}{
					OpeningDateKey:           filter.OpeningDate,
					ClosingDateKey:           date.NextDate(filter.OpeningDate),
					DistrictNameKey:          filter.DistrictName,
					AddressKey:               filter.Address,
					ManagementCompanyNameKey: filter.ManagementCompanyName,
					ServiceOrganizationKey:   filter.ServiceOrganizationName,
					UrgencyCategoryNameKey:   filter.UrgencyCategoryName,
				},
			},
			&numberAnomalies,
		)
		if err != nil {
			util.LogError(err)
			return nil, err
		}

		var numberNormal []NumberTuple

		err = service.Conn.CallTyped(NumberNormal,
			[]interface{}{
				map[string]interface{}{
					OpeningDateKey:           filter.OpeningDate,
					ClosingDateKey:           date.NextDate(filter.OpeningDate),
					DistrictNameKey:          filter.DistrictName,
					AddressKey:               filter.Address,
					ManagementCompanyNameKey: filter.ManagementCompanyName,
					ServiceOrganizationKey:   filter.ServiceOrganizationName,
					UrgencyCategoryNameKey:   filter.UrgencyCategoryName,
				},
			},
			&numberNormal,
		)
		if err != nil {
			util.LogError(err)
			return nil, err
		}

		var percent uint32

		if numberNormal[0].Value == 0 {
			percent = 0
		} else {
			percent = uint32((100 * numberAnomalies[0].Value) / (numberAnomalies[0].Value + numberNormal[0].Value))
		}

		points = append(points, Point{Percent: percent, Date: filter.OpeningDate})

		filter.OpeningDate = date.NextDate(filter.OpeningDate)
	}

	return &Statistic{Points: points}, nil
}

func (service *RequestRegistry) InsertAnomaly(info *AnomalyInsertInfo) error {
	client, err := srv.Dial(service.AddressRegistry)
	if err != nil {
		util.LogError(err)
	}

	response, err := client.Do(&srv.Request{Address: info.Address})
	if err != nil {
		util.LogError(err)
	} else {
		info.Latitude = response.Latitude
		info.Longitude = response.Longitude
	}

	_, err = service.Conn.Insert(service.Anomalies.Name,
		[]interface{}{
			info.Id,
			info.OpeningDate,
			info.ClosingDate,
			info.DistrictName,
			info.Address,
			info.FaultName,
			info.ManagementCompanyName,
			info.ServiceOrganizationName,
			info.UrgencyCategoryName,
			info.Latitude,
			info.Longitude,
		})
	if err != nil {
		util.LogError(err)
	}
	return nil
}

func (service *RequestRegistry) InsertNormal(info *NormalInsertInfo) error {
	_, err := service.Conn.Insert(service.Normal.Name, []interface{}{
		info.Id,
		info.OpeningDate,
		info.ClosingDate,
		info.DistrictName,
		info.Address,
		info.FaultName,
		info.ManagementCompanyName,
		info.ServiceOrganizationName,
		info.UrgencyCategoryName,
	})
	if err != nil {
		util.LogError(err)
	}
	return err
}

func (service *RequestRegistry) InsertRequest(info *RequestInsertInfo) error {
	_, err := service.Conn.Insert(service.Requests.Name, []interface{}{
		info.RequestRootIdentifier,
		info.OpeningDate,
		info.ClosingDate,
		info.DistrictName,
		info.Address,
		info.FaultName,
		info.ManagementCompanyName,
		info.ServiceOrganizationName,
		info.UrgencyCategoryName,
		info.AnomalyCategory,
		info.Effectiveness,
		info.Feedback,
		info.GroupId,
	})
	if err != nil {
		util.LogError(err)
	}
	return err
}
