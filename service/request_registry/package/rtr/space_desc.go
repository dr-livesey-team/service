package rtr

import (
	"encoding/json"
	"io/ioutil"
)

const (
	NumberType string = "number"
	StringType string = "string"
)

type FieldDesc struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type SpaceDesc struct {
	Name   string      `json:"name"`
	Fields []FieldDesc `json:"fields"`
}

func LoadSpaceDesc(filename string) (*SpaceDesc, error) {
	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return ParseSpaceDesc(buffer)
}

func ParseSpaceDesc(buffer []byte) (*SpaceDesc, error) {
	var desc SpaceDesc

	err := json.Unmarshal(buffer, &desc)

	if err != nil {
		return nil, err
	}

	return &desc, nil
}

