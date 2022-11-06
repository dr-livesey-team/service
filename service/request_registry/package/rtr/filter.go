package rtr

import (
	"encoding/json"
)

type Filter struct {
	OpeningDate  		      string `json:"opening date"`
	ClosingDate  		      string `json:"closing date"`
	DistrictName 		    []string `json:"district name"`
	Address      	        []string `json:"address"`
	ManagementCompanyName   []string `json:"management company name"`
	ServiceOrganizationName []string `json:"service organization name"`
	UrgencyCategoryName     []string `json:"urgency category name"`
	AnomalyCategory         []string `json:"anomaly category"`
}

func MarshalFilter(filter *Filter) ([]byte, error) {
	buffer, err := json.Marshal(filter)
	if err != nil {
		return nil, err
	}

	buffer = append(buffer, '\n')

	return buffer, nil
}

func UnmarshalFilter(buffer []byte) (*Filter, error) {
	var filter Filter

	err := json.Unmarshal(buffer, &filter)
	if err != nil {
		return nil, err
	}

	return &filter, nil
}
