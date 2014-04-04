package sockets

import (
	"log"
)

type SocketManager struct {
	Broadcast chan []byte

	Register   chan *SocketConnection
	Unregister chan *SocketConnection

	Connections map[*SocketConnection]bool
}

// There is a single SocketManager that handles all sockets
var Manager SocketManager

func init() {
	Manager = SocketManager{
		Broadcast:   make(chan []byte),

		Register:   make(chan *SocketConnection),
		Unregister: make(chan *SocketConnection),

		Connections: make(map[*SocketConnection]bool),
	}

	// Set up the slices at the keys we will need

	go Manager.AcceptConnections()
	go Manager.AcceptBroadcasts()
}

// Accepting incoming connections

func (m *SocketManager) RegisterConnection(conn *SocketConnection) {
	log.Println("New connection", conn)
	// Subscribe to top-level currency
	m.Connections[conn] = true
}

func (m *SocketManager) UnregisterConnection(conn *SocketConnection) {
	delete(m.Connections, conn)
	go conn.Connection.Close()
}

func (m *SocketManager) AcceptConnections() {
	for {
		select {
		case c := <-m.Register:
			m.RegisterConnection(c)
		case c := <-m.Unregister:
			m.UnregisterConnection(c)
		}
	}
}

func (m *SocketManager) AcceptBroadcasts() {
	for {
		select {
		case msg := <-m.Broadcast:
			log.Println(string(msg))
		}
	}
}



