package main

import (
	"log"
	"net/http"
	customwebsocket "chatapplication/websocket"
)


func serverWS(pool *customwebsocket.Pool, w http.ResponseWriter, r *http.Request) {
	log.Println("This is working")
	conn, err := customwebsocket.Upgrade(w, r)

	if err != nil {
		log.Println(err)
		return
	}

	client := &customwebsocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	log.Println("This is working")
	pool := customwebsocket.NewPool()

	// go routine
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serverWS(pool, w, r)
	})

}

func main() {

	setupRoutes()
	http.ListenAndServe(":9000", nil)
}