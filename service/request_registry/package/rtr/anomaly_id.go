package rtr

import (
	"encoding/json"
)

type AnomalyId struct {
	Id uint64 `json:"id"`
}

func MarshalAnomalyId(id *AnomalyId) ([]byte, error) {
	buffer, err := json.Marshal(id)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func UnmarshalAnomalyId(buffer []byte) (*AnomalyId, error) {
	var id AnomalyId

	err := json.Unmarshal(buffer, &id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

