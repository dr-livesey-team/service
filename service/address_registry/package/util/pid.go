package util

import (
	"fmt"
	"os"
)

func LogPid(filename string) error {
	pid := os.Getpid()

	buffer := fmt.Sprintf("%d", pid)

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	_, err = file.Write([]byte(buffer))
	if err != nil {
		return err
	}

	return nil
}