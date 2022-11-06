package rtr

import (
	"encoding/json"
)

type AnomalySelectInfo struct {
	Id 						uint64   `json:"id"`
	OpeningDate 			string   `json:"opening_date"`
	ClosingDate 			string   `json:"closing_date"`
    DistrictName 			string   `json:"district_name"`
	Address      	        string   `json:"address"`
	FaultName    		    string   `json:"fault_name"`
	ManagementCompanyName   string   `json:"management_company_name"`
	ServiceOrganizationName string   `json:"service_organization_name"`
	UrgencyCategoryName     string   `json:"urgency_category_name"`
	AnomalyCategory         string   `json:"anomaly_category"`
	Latitude  				float64  `json:"latitude"`
	Longitude 				float64  `json:"longitude"`
}

func MarshalAnomalySelectInfo(info *AnomalySelectInfo) ([]byte, error) {
	buffer, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func UnmarshalAnomalySelectInfo(buffer []byte) (*AnomalySelectInfo, error) {
	var info AnomalySelectInfo

	err := json.Unmarshal(buffer, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

type AnomalySelectInfos struct {
	Infos []AnomalySelectInfo `json:"anomalies"`
}

func MarshalAnomalySelectInfos(infos *AnomalySelectInfos) ([]byte, error) {
	buffer, err := json.Marshal(infos)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func UnmarshalAnomalySelectInfos(buffer []byte) (*AnomalySelectInfos, error) {
	var infos AnomalySelectInfos

	err := json.Unmarshal(buffer, &infos)
	if err != nil {
		return nil, err
	}

	return &infos, nil
}
