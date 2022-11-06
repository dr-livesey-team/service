package rtr

import (
	"encoding/json"
)

type RequestSelectInfo struct {
	RequestRootIdentifier     uint64 `json:"request_root_identifier"`
	OpeningDate  		      string `json:"opening_date"`
	ClosingDate  		      string `json:"closing_date"`
	DistrictName 		      string `json:"district_name"`
	Address      	          string `json:"address"`
	FaultName    		      string `json:"fault_name"`
	ManagementCompanyName     string `json:"management_company_name"`
	ServiceOrganizationName   string `json:"service_organization_name"`
	UrgencyCategoryName       string `json:"urgency_category_name"`
	Feedback                  string `json:"feedback"`
	GroupId					  uint64 `json:"group_id"`
}

func MarshalRequestSelectInfo(info *RequestSelectInfo) ([]byte, error) {
	buffer, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func UnmarshalRequestSelectInfo(buffer []byte) (*RequestSelectInfo, error) {
	var info RequestSelectInfo

	err := json.Unmarshal(buffer, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

type RequestSelectInfos struct {
	Infos []RequestSelectInfo `json:"requests"`
}

func MarshalRequestSelectInfos(infos *RequestSelectInfos) ([]byte, error) {
	buffer, err := json.Marshal(infos)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func UnmarshalRequestSelectInfos(buffer []byte) (*RequestSelectInfos, error) {
	var infos RequestSelectInfos

	err := json.Unmarshal(buffer, &infos)
	if err != nil {
		return nil, err
	}

	return &infos, nil
}