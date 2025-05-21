package customwebsocket


import (
	"log"
	// "fmt"
	"net/http"
	"github.com/gorilla/websocket"
)
var upgrader = websocket.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true}

	connection , error := upgrader.Upgrade(w, r, nil)

	if error != nil {
		log.Println("Web socket Connection error", error)
	}
	return connection, nil
}