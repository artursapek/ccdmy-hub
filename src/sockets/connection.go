package sockets

import (
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type SocketConnection struct {
	// Meta info
	OpenTime   int64
	RemoteAddr string

	Username string

	// Data transfer
	Connection *websocket.Conn
	Send       chan []byte
}

func NewSocketConnection(ws *websocket.Conn, remoteAddr string) *SocketConnection {
	return &SocketConnection{
		OpenTime:   time.Now().Unix(),
		RemoteAddr: remoteAddr,
		Send:       make(chan []byte, 256),
		Connection: ws,
	}
}

func (c *SocketConnection) Writer() {
	for msg := range c.Send {
		err := c.Connection.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			break
		}
	}
	c.Connection.Close()
}

func (c *SocketConnection) Reader() {
	for {
		_, message, err := c.Connection.ReadMessage()
		if err != nil {
			break
		}
		log.Println(string(message))
	}
	c.Connection.Close()
}

