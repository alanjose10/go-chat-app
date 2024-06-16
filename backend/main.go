package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alanjose10/go-chat-app/pkg/websocket"
)

func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Println(err)
	}
	go websocket.Writer(ws)
	websocket.Reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Chat App v0.2")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
