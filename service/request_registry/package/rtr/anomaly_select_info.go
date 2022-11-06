package rtr

import (
	"encoding/json"
)

type AnomalySelectInfo struct {
	Id        uint64  `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
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
