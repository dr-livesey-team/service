package rtr

import "encoding/json"

type Point struct {
	Date    string  `json:"date"`
	Percent uint32 `json:"percent"`
}

type Statistic struct {
	Points []Point `json:"points"`
}

func MarshalStatistic(statistic *Statistic) ([]byte, error) {
	buffer, err := json.Marshal(statistic)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func UnmarshalStatistic(buffer []byte) (*Statistic, error) {
	var statistic Statistic

	err := json.Unmarshal(buffer, &statistic)
	if err != nil {
		return nil, err
	}

	return &statistic, nil
}
