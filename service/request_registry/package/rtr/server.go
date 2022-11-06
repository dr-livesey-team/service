package rtr

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/elisfromkirov/service/service/request_registry/package/util"
)

type Server struct {
	Listener net.Listener
	Conns    chan net.Conn
	Signals  chan os.Signal
	Service  *RequestRegistry
}

func StartServer(endpoint string, service *RequestRegistry) (*Server, error) {
	listener, err := net.Listen("tcp", endpoint)
	if err != nil {
		return nil, err
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT)

	conns := make(chan net.Conn, 1)

	return &Server{listener, conns, signals, service}, nil
}

func (server *Server) Stop() error {
	return server.Listener.Close()
}

func (server *Server) Listen() {
	for {
		conn, err := server.Listener.Accept()
		if err != nil {
			util.LogError(err)
		}

		server.Conns <- conn
	}
}

func (server *Server) ListenAndServe() {
	go server.Listen()
	for {
		select {
		case conn := <- server.Conns:
			handler := NewHandler(server.Service)
			go handler.Handle(conn)
		case <- server.Signals:
			return
		}
	}
}
