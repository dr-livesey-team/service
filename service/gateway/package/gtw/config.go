package gtw

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Listen 					string `json:"listen"`
	PidFile 			    string `json:"pid"`
	LogFile 			    string `json:"log"`
	RequestRegistryEndpoint string `json:"rtr"`
}

func LoadConfig(filename string) (*Config, error) {
	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return ParseConfig(buffer)
}

func ParseConfig(buffer []byte) (*Config, error) {
	var config Config

	err := json.Unmarshal(buffer, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
