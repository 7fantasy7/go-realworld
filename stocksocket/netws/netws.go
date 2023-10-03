package netws

import (
	"fmt"
	"golang.org/x/net/websocket"
	"io"
	"net/http"
)

type Server struct {
	clients map[*websocket.Conn]bool
}

// To store the WebSocket connection immediately after the upgrade request,
// you would need to use another WebSocket library or package in Go,
// since the golang.org/x/net/websocket package combines
// the upgrade and connection handling into a single step (via the Handler function),
// and does not provide control immediately after the upgrade request.
func (server *Server) echo(ws *websocket.Conn) {
	fmt.Println("New connection")

	io.Copy(ws, ws)

	fmt.Println("Connection closed")
}

func StartServer() *Server {
	server := Server{
		make(map[*websocket.Conn]bool),
	}

	http.Handle("/", websocket.Handler(server.echo))
	go http.ListenAndServe(":8080", nil)

	return &server
}
