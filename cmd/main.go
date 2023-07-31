package main

import (
	"fmt"
	"log"
	"net/http"
	"sevens/internal/serverLib"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func serveWebSocket(server *serverLib.Server, w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade websocket", err)
		return
	}
	user := serverLib.NewUser(server, connection)
	go user.HandleWrite()
	go user.HandleRead()
	server.Register <- user
}

func main() {
	server := serverLib.NewServer()
	go server.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWebSocket(server, w, r)
	})

	port := 8080
	log.Printf("Server started on port %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
