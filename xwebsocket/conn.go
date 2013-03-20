package xwebsocket

import (
    //"sync"
  //  "time"
   // "fmt"
   // "net/http"
)
import (
   "code.google.com/p/go.net/websocket"
)

type WsConnection struct {
	// The websocket connection.
	ws *websocket.Conn
 
	// Buffered channel of outbound messages.
	send chan string
}

func (c *WsConnection) Reader() {
	for {
		var message string
		err := websocket.Message.Receive(c.ws, &message)
		if err != nil {
			break
		}
		Hubb.broadcast <- message
	}
	c.ws.Close()
}
 
func (c *WsConnection) Writer() {
	for message := range c.send {
		err := websocket.Message.Send(c.ws, message)
		if err != nil {
			break
		}
	}
	c.ws.Close()
}
 
func WsHandler(ws *websocket.Conn) {
	c := &WsConnection{send: make(chan string, 256), ws: ws}
	Hubb.register <- c
	defer func() { Hubb.unregister <- c }()
	go c.Writer()
	c.Reader()
}

