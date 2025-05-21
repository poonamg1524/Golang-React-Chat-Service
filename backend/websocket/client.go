package customwebsocket

import (
	// "net/http"
	"github.com/gorilla/websocket"
	"sync"
	"fmt"
)

type Client struct {
	Conn *websocket.Conn //Websocket Connection
	Pool *Pool
	mu   sync.Mutex
}

type Message struct {
	Type int `json:"type"`
	Body string `json:"body"`
}


func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		//Read each msg
		msgType, msg, err := c.Conn.ReadMessage()

		if err != nil {
			fmt.Println("Error while reading the message", err)
			return
		}
		
		//Create a broadcast of an msg
		m := Message{Type: msgType, Body: string(msg)}

		//sending msg to broadcast Channel
		c.Pool.Broadcast <- m 

		fmt.Println("Msg received ==>", msg)
	}
}