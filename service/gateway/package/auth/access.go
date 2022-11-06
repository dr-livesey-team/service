package auth

import (
	"encoding/json"
)

type Access struct {
	Access bool `json:"access"`
}

func MarshalAccess(info *Access) ([]byte, error) {
	buffer, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}
