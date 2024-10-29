package main

import (
	"backend/pkg/websocket"
	
	"fmt"
	"log"
	"net/http"
)

// WebSocket enpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	// Prints out the request host
	fmt.Println(r.Host)

	// Upgrades the connection to a WebSocket connection
	ws, err := websocket.Upgrade(w, r)

	// Print out the error message if there is one
	if err != nil {
		log.Println(err)
	}

	go websocket.Writer(ws)
	websocket.Reader(ws)
}

func setupRoutes() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Simple Server")
	// })

	// map the /ws endpoint to the serveWs function
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Distributed Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}