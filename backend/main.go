package main

import (
	"backend/pkg/websocket"

	"fmt"
	"net/http"
)

// WebSocket enpoint
func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	// Prints out the request host
	fmt.Println(r.Host)
	fmt.Println("WebSocket Endpoint Hit")

	// Upgrades the connection to a WebSocket connection
	conn, err := websocket.Upgrade(w, r)

	// Print out the error message if there is one
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	// Create a new client on every connection that is registered with a pool.
    client := &websocket.Client{
        Conn: conn,
        Pool: pool,
    }

    pool.Register <- client
    client.Read()
}

// Creates a new pool to provide to the websocket.
func setupRoutes() {
    pool := websocket.NewPool()
    go pool.Start()

    http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        serveWs(pool, w, r)
    })
}

func main() {
	fmt.Println("Distributed Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}