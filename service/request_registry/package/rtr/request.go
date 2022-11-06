package rtr

import (
	"encoding/json"
)

type Request struct {
	Function string `json:"func"`
	Buffer   string `json:"data"`
}

func MarshalRequest(request *Request) ([]byte, error) {
	buffer, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func UnmarshalRequest(buffer []byte) (*Request, error) {
	var request Request

	err := json.Unmarshal(buffer, &request)
	if err != nil {
		return nil, err
	}

	return &request, nil
}