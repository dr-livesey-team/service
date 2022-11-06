package srv

import (
	"encoding/json"
)

type Request struct {
	Address string `json:"address"`
}

func MarshalRequest(request *Request) ([]byte, error) {
	buffer, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	buffer = append(buffer, '\n')

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
