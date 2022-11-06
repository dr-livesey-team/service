package rtr

import (
	"encoding/json"
)

type Response struct {
	Function string `json:"func"`
	Buffer   string `json:"data"`
}

func MarshalResponse(response *Response) ([]byte, error) {
	buffer, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func UnmarshalResponse(buffer []byte) (*Response, error) {
	var response Response

	err := json.Unmarshal(buffer, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

