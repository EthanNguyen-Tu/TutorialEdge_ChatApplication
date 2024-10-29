package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
    ID   string // a uniquely identifiably string for a particular connection
    Conn *websocket.Conn // a pointer to a websocket.Conn object
    Pool *Pool // a pointer to the Pool which the client is a part of
}

type Message struct {
    Type int    `json:"type"`
    Body string `json:"body"`
}

// Constantly listens in for new messages coming through on the Client’s
// websocket connection, passing any messages that come in to the pool's
// Broadcast channel.
func (c *Client) Read() {
    defer func() {
        c.Pool.Unregister <- c
        c.Conn.Close()
    }()

    for {
        messageType, p, err := c.Conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }
        message := Message{Type: messageType, Body: string(p)}

		// Broadcast's the recieved message to every client within the pool.
        c.Pool.Broadcast <- message
		
        fmt.Printf("Message Received: %+v\n", message)
    }
}