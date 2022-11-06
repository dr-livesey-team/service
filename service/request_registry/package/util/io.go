package util

import (
	"bufio"
	"net"
)

func ReadMessage(conn net.Conn) ([]byte, error) {
	reader := bufio.NewReader(conn)

	buffer, err := reader.ReadBytes('\n')
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func WriteMessage(conn net.Conn, buffer []byte) error {
	writer := bufio.NewWriter(conn)

	_, err := writer.Write(buffer)
	if err != nil {
		return err
	}

	err = writer.WriteByte('\n')
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return nil
	}

	return nil
}
