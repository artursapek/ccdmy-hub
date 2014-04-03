package main

import (
	"log"
	"github.com/gorilla/websocket"
	"net/http"
	"sockets"
)

func LiveTickerHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)

	if _, ok := err.(websocket.HandshakeError); ok {
	  http.Error(w, "Not a websocket handshake", 400)
	  return
	} else if err != nil {
	  log.Println("Unknown error:")
	  log.Println(err)
	  return
	}

	c := sockets.NewSocketConnection(ws, r.RemoteAddr)

	defer func () { sockets.Manager.Unregister <- c }()

	c.Writer()
}

func main() {
	log.Println("Accepting connections at :8080")
	http.HandleFunc("/", LiveTickerHandler)
	http.ListenAndServe(":8080", nil)
}
