package srv

import (
	"bufio"
	"net"
)

type Client struct {
	Conn net.Conn
}

func Dial(endpoint string) (*Client, error) {
	conn, err := net.Dial("tcp", endpoint)
	if err != nil {
		return nil, err
	}

	return &Client{Conn: conn}, nil
}

func (client* Client) Do(request *Request) (*Response, error) {
	buffer, err := MarshalRequest(request)
	if err != nil {
		return nil, err
	}

	err = WriteRequest(client.Conn, buffer)
	if err != nil {
		return nil, err
	}

	buffer, err = ReadResponse(client.Conn)
	if err != nil {
		return nil, err  
	}

	response, err := UnmarshalResponse(buffer)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (client *Client) Close() error {
	return client.Conn.Close()
}

func WriteRequest(conn net.Conn, buffer []byte) error {
	writer := bufio.NewWriter(conn)
	
	_, err := writer.Write(buffer)
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return nil
	}

	return nil
}

func ReadResponse(conn net.Conn) ([]byte, error) {
	reader := bufio.NewReader(conn)

	buffer, err := reader.ReadBytes('\n')
	if err != nil {
		return nil, err
	}

	return buffer, nil
}