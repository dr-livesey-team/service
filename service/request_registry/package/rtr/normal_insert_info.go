package rtr

import (
	"encoding/json"
)

type NormalInsertInfo struct {
	Id uint64 `json:"id"`
	OpeningDate string `json:"opening_date"`
	ClosingDate string `json:"closing_date"`
	DistrictName string `json:"district_name"`
	Address string `json:"address"`
	FaultName string `json:"fault_name"`
	ManagementCompanyName string `json:"management_company_name"`
	ServiceOrganizationName string `json:"service_organization_name"`
	UrgencyCategoryName string `json:"urgency_category_name"`
}

func MarshalNormalInsertInfo(info *NormalInsertInfo) ([]byte, error) {
	buffer, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func UnmarshalNormalInsertInfo(buffer []byte) (*NormalInsertInfo, error) {
	var info NormalInsertInfo

	err := json.Unmarshal(buffer, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

