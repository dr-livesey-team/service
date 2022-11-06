package rtr

import (
	"encoding/json"
)

type RequestInsertInfo struct {
	RequestRootIdentifier uint64 `json:"request_root_identifier"`
	OpeningDate string `json:"opening_date"`
	ClosingDate string `json:"closing_date"`
	DistrictName string `json:"district_name"`
	Address string `json:"address"`
	FaultName string `json:"fault_name"`
	ManagementCompanyName string `json:"management_company_name"`
	ServiceOrganizationName string `json:"service_organization_name"`
	UrgencyCategoryName string `json:"urgency_category_name"`
	AnomalyCategory uint64 `json:"anomaly_category"`
	Effectiveness string `json:"effectiveness"`
	Feedback string `json:"feedback"`
	GroupId uint64 `json:"group_id"`
}

func MarshalRequestInsertInfo(info *RequestInsertInfo) ([]byte, error) {
	buffer, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func UnmarshalRequestInsertInfo(buffer []byte) (*RequestInsertInfo, error) {
	var info RequestInsertInfo

	err := json.Unmarshal(buffer, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
