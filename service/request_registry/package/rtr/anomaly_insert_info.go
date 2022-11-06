package rtr

import (
	"encoding/json"
)

type AnomalyInsertInfo struct {
	Id uint64 `json:"id"`
	OpeningDate	string `json:"opening_date"`
	ClosingDate	string `json:"closing_date"`
    DistrictName string `json:"district_name"`
	Address string `json:"address"`
	FaultName string `json:"fault_name"`
	ManagementCompanyName string `json:"management_company_name"`
	ServiceOrganizationName string `json:"service_organization_name"`
	UrgencyCategoryName string `json:"urgency_category_name"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func MarshalAnomalyInsertInfo(info *AnomalyInsertInfo) ([]byte, error) {
	buffer, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func UnmarshalAnomalyInsertInfo(buffer []byte) (*AnomalyInsertInfo, error) {
	var info AnomalyInsertInfo

	err := json.Unmarshal(buffer, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
