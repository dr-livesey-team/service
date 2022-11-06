package rtr

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Listen  		  string `json:"listen"`
	PidFile 		  string `json:"pid"`
	LogFile 		  string `json:"log"`
	DataBase          string `json:"database"`
	User    		  string `json:"user"`
	Pass    	      string `json:"pass"`
	RequestsDescFile  string `json:"requests"`
	AnomaliesDescFile string `json:"anomalies"`
	NormalDescFile 	  string `json:"normal"`
	AddressRegistry   string `json:"adr"`
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
