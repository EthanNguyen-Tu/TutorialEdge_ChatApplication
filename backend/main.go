package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,

	// Checks the origin of the connection to allow making requests
	// the React development server
	// NOTE: Currently does no checking, allowing any connection
	CheckOrigin: func(r *http.Request) bool { return true },
}

// listens for new messages being sent to the WebScoket endpoint
func reader(conn *websocket.Conn) {
	for {
		// Reads in a message p
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// Prints out the read message
		fmt.Println(string(p))
		
		// Print out the error message if there is one
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

// WebSocket enpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	// Prints out the request host
	fmt.Println(r.Host)

	// Upgrades the connection to a WebSocket connection
	ws, err := upgrader.Upgrade(w, r, nil)

	// Print out the error message if there is one
	if err != nil {
		log.Println(err)
	}

	// List indefinitely for new messages coming through
	// the WebSocket connection
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})

	// map the /ws endpoint to the serveWs function
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}