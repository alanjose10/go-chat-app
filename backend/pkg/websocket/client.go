package websocket

import (
	"log"

	"github.com/gorilla/websocket"
)

type Message struct {
	messageType int    `json:"type"`
	messageBody string `json:"body"`
}

type Pool struct {
	register   chan *Client
	unRegister chan *Client
	clients    map[*Client]bool
	broadcast  chan Message
}

func (pool *Pool) Start() {
	for {
		select {}
	}
}

func NewPool() *Pool {
	return &Pool{
		register:   make(chan *Client),
		unRegister: make(chan *Client),
		clients:    make(map[*Client]bool),
		broadcast:  make(chan Message),
	}
}

type Client struct {
	id   string
	conn *websocket.Conn
	pool *Pool
}

func (c *Client) Read() {
	defer func() {
		c.pool.unRegister <- c
		c.conn.Close()
	}()

	for {
		messageType, p, err := c.conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		message := Message{messageType: messageType, messageBody: string(p)}
		c.pool.broadcast <- message
		log.Printf("Message received: %+v\n", message)
	}
}
