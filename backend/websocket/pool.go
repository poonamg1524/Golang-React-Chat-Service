package customwebsocket

import (
	// "log"
	"fmt"
	// "net/http"
	// "github.com/gorilla/websocket"
)

type Pool struct {
	Register chan *Client
	Unregister chan *Client
	Clients map[*Client]bool //map -> tell if a channel is active or not
	Broadcast chan Message
}

//Return New Pool with new Memory location
func NewPool() *Pool {

	return &Pool{
		Register:	make(chan *Client),
		Unregister:	make(chan *Client),
		Clients:	make(map[*Client]bool),
		Broadcast:	make(chan Message),
	}
}


func (pool *Pool) Start() {
	for{
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("total Connection pool:=", len(pool.Clients))
			for k, _ := range pool.Clients {
				fmt.Println(k)
				k.Conn.WriteJSON( Message{
										Type: 1, 
										Body: "User Disconnected"})	
			}
			break


		case client := <-pool.Unregister:
			//We need to delete the client value from map
			delete(pool.Clients, client)
			fmt.Println("Current total of Connection pool:-", len(pool.Clients))
			for k, _ := range pool.Clients {
				fmt.Println(k)
				k.Conn.WriteJSON( Message{
										Type: 1, 
										Body: "New user Added"})	
			}
			break

			
		case msg := <-pool.Broadcast:
			fmt.Println("Broadcasting a message")

			for k, _ := range pool.Clients {
				if err := k.Conn.WriteJSON(msg); err != nil {
					fmt.Println("Error while broadcasting the message", err)
				}
			}
		}
	}
}