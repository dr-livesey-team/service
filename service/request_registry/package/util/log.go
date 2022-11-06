package util

import (
	"errors"
	"fmt"
	"os"
)

const (
	Debug   int = 0
	Warning int = 1
	Error   int = 2
)

var (
	file *os.File
)

func OpenLogFile(name string) error {
	var err error

	file, err = os.Create(name)
	return err
}

func CloseLogFile() error {
	return file.Close()
}

func Log(severity int, format string, args ...any) {
	err := WriteLogHeader(severity)
	if err != nil {
		panic(err)
	}

	err = WriteLogBody(format, args...)
	if err != nil {
		panic(err)
	}

	err = WriteLogFooter()
	if err != nil {
		panic(err)
	}

	err = file.Sync()
	if err != nil {
		panic(err)
	}
}

func WriteLogHeader(severity int) error {
	var err error

	switch severity {
	case Debug:
		_, err = file.Write([]byte("[ DEBUG ] "))
	case Warning:
		_, err = file.Write([]byte("[WARNING] "))
	case Error:
		_, err = file.Write([]byte("[ ERROR ] "))
	default:
		err = errors.New("unknown log severity")
	}
	return err
}

func WriteLogBody(format string, args ...any) error {
	_, err := file.Write([]byte(fmt.Sprintf(format, args...)))
	return err
}

func WriteLogFooter() error {
	_, err := file.Write([]byte("\n"))
	return err
}

func LogError(err error) {
	Log(Error, "%s", err.Error())
}
