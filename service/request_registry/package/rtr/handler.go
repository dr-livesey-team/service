package rtr

import (
	"errors"
	"github.com/elisfromkirov/service/service/request_registry/package/util"
	"net"
)

type Handler struct {
	Service *RequestRegistry
}

func NewHandler(service *RequestRegistry) *Handler {
	return &Handler{Service:
		service,
	}
}

func (handler *Handler) Handle(conn net.Conn) {
	defer conn.Close()

	for {
		buffer, err := util.ReadMessage(conn)
		if err != nil {
			util.LogError(err)
			return
		}
		util.Log(util.Debug, "Read message: %s", buffer)

		request, err := UnmarshalRequest(buffer)
		if err != nil {
			util.LogError(err)
			return
		}

		response, err := handler.Call(request)
		if err != nil {
			util.LogError(err)
			return
		}

		buffer, err = MarshalResponse(response)
		if err != nil {
			util.LogError(err)
			return
		}

		err = util.WriteMessage(conn, buffer)
		if err != nil {
			util.LogError(err)
		}
	}
}

func (handler *Handler) Call(request *Request) (*Response, error) {
	switch request.Function {
	case GetAnomaliesFunction:
		filter, err := UnmarshalFilter([]byte(request.Buffer))
		if err != nil {
			return nil, err
		}

		infos, err := handler.Service.GetAnomalies(filter)
		if err != nil {
			return nil, err
		}

		buffer, err := MarshalAnomalySelectInfos(infos)
		if err != nil {
			return nil, err
		}

		return &Response{Function: GetAnomaliesFunction, Buffer: string(buffer)}, nil
	case GetRequestsFunction:
		id, err := UnmarshalAnomalyId([]byte(request.Buffer))
		if err != nil {
			return nil, err
		}

		infos, err := handler.Service.GetRequests(id)
		if err != nil {
			return nil, err
		}

		buffer, err := MarshalRequestSelectInfos(infos)
		if err != nil {
			return nil, err
		}

		return &Response{Function: GetRequestsFunction, Buffer: string(buffer)}, nil
	case GetStatisticFunction:
		filter, err := UnmarshalFilter([]byte(request.Buffer))
		if err != nil {
			return nil, err
		}

		statistic, err := handler.Service.GetStatistic(filter)
		if err != nil {
			return nil, err
		}

		buffer, err := MarshalStatistic(statistic)
		if err != nil {
			return nil, err
		}

		return &Response{Function: GetAnomaliesFunction, Buffer: string(buffer)}, nil

	case InsertAnomalyFunction:
		info, err := UnmarshalAnomalyInsertInfo([]byte(request.Buffer))
		if err != nil {
			return nil, err
		}

		err = handler.Service.InsertAnomaly(info)
		if err != nil {
			return nil, err
		}

		return &Response{Function: GetAnomaliesFunction, Buffer: ""}, nil
	case InsertNormalFunction:
		info, err := UnmarshalNormalInsertInfo([]byte(request.Buffer))
		if err != nil {
			return nil, err
		}

		err = handler.Service.InsertNormal(info)
		if err != nil {
			return nil, err
		}

		return &Response{Function: GetAnomaliesFunction, Buffer: ""}, nil
	case InsertRequestFunction:
		info, err := UnmarshalRequestInsertInfo([]byte(request.Buffer))
		if err != nil {
			return nil, err
		}

		err = handler.Service.InsertRequest(info)
		if err != nil {
			return nil, err
		}

		return &Response{Function: InsertRequestFunction, Buffer: ""}, nil
	}
	return nil, errors.New("call unknown function")
}
