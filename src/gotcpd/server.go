package gotcpd

import (
	"net"
	"fmt"
	"log"
)

type Server struct {
	Host              string
	Port              uint16

	listener          net.Listener
	connectionChannel chan net.Conn
	stopChannel       chan interface{}
}

func NewServer(host string, port uint16) Server {
	return Server{
		Host: host,
		Port: port,

		connectionChannel: make(chan net.Conn),
		stopChannel: make(chan interface{}),
	}
}

func (server Server) Listen() error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Host, server.Port))

	if err != nil {
		return err
	}

	server.listener = listener

	go server.acceptConnections()

	return nil
}

func (server Server) Stop() error {
	close(server.stopChannel)

	return server.listener.Close()
}

func (server Server) GetConnectionChannel() chan net.Conn {
	return server.connectionChannel
}

func (server Server) acceptConnections() {
	for {
		select {
		case <-server.stopChannel:
			return
		default:
		}

		connection, err := server.listener.Accept()

		if err != nil {
			log.Printf("Cannot accept connection: %s", err)

			continue
		}

		server.connectionChannel <- connection
	}
}
