package gotcpd

import (
	"net"
	"log"
	"bytes"
)

type ConnectionHandler struct {
	channel     chan net.Conn
	stopChannel chan interface{}
}

func NewConnectionHandler(channel chan net.Conn) ConnectionHandler {
	return ConnectionHandler{
		channel: channel,
		stopChannel: make(chan interface{}),
	}
}

func (connectionHandler ConnectionHandler) Run() {
	go connectionHandler.worker()
}

func (connectionHandler ConnectionHandler) Stop() {
	close(connectionHandler.stopChannel)
}

func (connectionHandler ConnectionHandler) worker() {
	for {
		select {
		case <-connectionHandler.stopChannel:
			return
		case connection := <-connectionHandler.channel:
			go connectionHandler.handle(connection)
		}
	}
}

func (connectionHandler ConnectionHandler) handle(connection net.Conn) {
	log.Printf("New connection: %s", connection)

	for {
		select {
		case <-connectionHandler.stopChannel:
			return
		default:
		}

		readBuffer := &bytes.Buffer{}
		_, err := readBuffer.ReadFrom(connection)

		if err != nil {
			log.Printf("Cannot read from connection: %s", err)

			return
		}

		log.Printf("Read: %s", readBuffer.String())
	}

	log.Printf("Closing connection: %s", connection)

	connection.Close()
}
