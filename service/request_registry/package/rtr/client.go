package rtr

import (
	"github.com/elisfromkirov/service/service/request_registry/package/util"
	"net"
)

const (
	GetAnomaliesFunction string = "GetAnomalies"
	GetRequestsFunction  string = "GetRequests"
	GetStatisticFunction string = "GetStatistic"
	InsertAnomalyFunction string = "InsertAnomaly"
	InsertNormalFunction string = "InsertNormal"
	InsertRequestFunction string = "InsertRequest"
)

type Client struct {
	Conn net.Conn
	Service *RequestRegistry
}

func Dial(endpoint string) (*Client, error) {
	conn, err := net.Dial("tcp", endpoint)
	if err != nil {
		return nil, err
	}

	return &Client{Conn: conn}, nil
}

func (client *Client) GetAnomalies(filter *Filter) (*AnomalySelectInfos, error) {
	buffer, err := MarshalFilter(filter)
	if err != nil {
		return nil, err
	}

	buffer, err = client.RemoteCall(GetAnomaliesFunction, buffer)
	if err != nil {
		return nil, err
	}

	infos, err := UnmarshalAnomalySelectInfos(buffer)
	if err != nil {
		return nil, err
	}

	return infos, err
}

func (client *Client) GetRequests(id *AnomalyId) (*RequestSelectInfos, error) {
	buffer, err := MarshalAnomalyId(id)
	if err != nil {
		return nil, err
	}

	buffer, err = client.RemoteCall(GetRequestsFunction, buffer)
	if err != nil {
		return nil, err
	}

	infos, err := UnmarshalRequestSelectInfos(buffer)
	if err != nil {
		return nil, err
	}

	return infos, err
}

func (client *Client) GetStatistic(filter *Filter) (*Statistic, error) {
	buffer, err := MarshalFilter(filter)
	if err != nil {
		return nil, err
	}

	buffer, err = client.RemoteCall(GetStatisticFunction, buffer)
	if err != nil {
		return nil, err
	}

	statistic, err := UnmarshalStatistic(buffer)
	if err != nil {
		return nil, err
	}

	return statistic, nil
}

func (client *Client) RemoteCall(function string, buffer []byte) ([]byte, error) {
	request := Request{Function: function, Buffer: string(buffer)}
	buffer, err := MarshalRequest(&request)
	if err != nil {
		return nil, err
	}

	err = util.WriteMessage(client.Conn, buffer)
	if err != nil {
		return nil, err
	}

	buffer, err = util.ReadMessage(client.Conn)
	if err != nil {
		return nil, err
	}

	response, err := UnmarshalResponse(buffer)
	if err != nil {
		return nil, err
	}

	return []byte(response.Buffer), nil
}