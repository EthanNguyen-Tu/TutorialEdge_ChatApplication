package websocket

import "fmt"

type Pool struct {
	Register   chan *Client // Will send out "New User Joined..."" to all clients within the pool when a new client connects.
	Unregister chan *Client // Will unregister a user and notify the pool when a client disconnects.
	Clients    map[*Client]bool // A map of clients to a boolean value used to dictate active/inactive, but not disconnected, based on browser focus.
	Broadcast  chan Message // A channel which, when passed a message, will loop through all clients in the pool and send the message through the socket connection.
}

// Creates a new Pool structure that contains all of the channels needed for concurrent communication and a map of clients.
func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

// Loops through all clients in the pool and sends a message depending on the type of event that happens.
func (pool *Pool) Start() {
    for {
        select {
        case client := <-pool.Register:
            pool.Clients[client] = true
            fmt.Println("Size of Connection Pool: ", len(pool.Clients))
            for client, _ := range pool.Clients {
                fmt.Println(client)
                client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})
            }
            break
        case client := <-pool.Unregister:
            delete(pool.Clients, client)
            fmt.Println("Size of Connection Pool: ", len(pool.Clients))
            for client, _ := range pool.Clients {
                client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
            }
            break
        case message := <-pool.Broadcast:
            fmt.Println("Sending message to all clients in Pool")
            for client, _ := range pool.Clients {
                if err := client.Conn.WriteJSON(message); err != nil {
                    fmt.Println(err)
                    return
                }
            }
        }
    }
}