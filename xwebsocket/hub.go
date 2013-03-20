package xwebsocket

import (
   // "sync"
)

type Hub struct {
	// Registered connections.
	connections map[*WsConnection]bool
 
	// Inbound messages from the connections.
	broadcast chan string
 
	// Register requests from the connections.
	register chan *WsConnection
 
	// Unregister requests from connections.
	unregister chan *WsConnection
}

func MakeNewHub() Hub{
	return Hub{
		broadcast:   make(chan string),
		register:    make(chan *WsConnection),
		unregister:  make(chan *WsConnection),
		connections: make(map[*WsConnection]bool),
	}
}

/*
var _init_ctxx sync.Once 
var GHub *Hub




func GetHub() *Hub {
	_init_ctxx.Do( func () { 
		
		GHub = Hub{
				broadcast:   make(chan string),
				register:    make(chan *WsConnection),
				unregister:  make(chan *WsConnection),
				connections: make(map[*WsConnection]bool)
		}
	})
    return GHub
}
*/
 
var Hubb = Hub{
	broadcast:   make(chan string),
	register:    make(chan *WsConnection),
	unregister:  make(chan *WsConnection),
	connections: make(map[*WsConnection]bool),
}
func DEADGetHubb() Hub{
	return Hubb
}


func (h *Hub) Run() {
	for {
		select {
		case c := <-h.register:
			h.connections[c] = true
		case c := <-h.unregister:
			delete(h.connections, c)
			close(c.send)
		case m := <-h.broadcast:
			for c := range h.connections {
				select {
				case c.send <- m:
				default:
					delete(h.connections, c)
					close(c.send)
					go c.ws.Close()
				}
			}
		}
	}
}