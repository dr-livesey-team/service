package srv

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/elisfromkirov/service/service/address_registry/package/util"
)

type Service struct {
	Listener net.Listener
	Conns chan net.Conn
	Signals chan os.Signal
}

func StartService(config *Config) (*Service, error) {
	listener, err := net.Listen("tcp", config.Listen)
	if err != nil {
		return nil, err
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT)

	conns := make(chan net.Conn, 1)

	return &Service{listener, conns, signals}, nil
}

func (service *Service) Stop() {
	err := service.Listener.Close()
	if err != nil {
		util.LogError(err)
	}
}

func (service *Service) Listen() {
	for {
		conn, err := service.Listener.Accept()
		if err != nil {
			util.LogError(err)
		}

		service.Conns <- conn
	}
}

func (service *Service) ListenAndServe() {
	go service.Listen()
	for {
		select {
		case conn := <- service.Conns:
			go Handle(conn)
		case <- service.Signals:
			return
		}
	}
}