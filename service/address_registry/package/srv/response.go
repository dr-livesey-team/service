package srv

import (
	"encoding/json"
)

type Response struct {
	Latitude float64 `json:"lattitude"`
	Longitude float64 `json:"longitude"`
}

func MarshalResponse(response *Response) ([]byte, error) {
	buffer, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}

	buffer = append(buffer, '\n')

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
