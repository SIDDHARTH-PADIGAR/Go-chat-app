package main

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn     *websocket.Conn
	username string
}

type Message struct {
	Type    string `json:"type"`
	Content string `json:"content"`
	Sender  string `json:"sender"`
	Room    string `json:"room"`
}

var clients = make(map[*Client]bool)
var clientsMutex = sync.Mutex{}

func handleClient(conn *websocket.Conn) {
	client := &Client{
		conn: conn,
	}

	// Wait for username message
	_, msg, err := conn.ReadMessage()
	if err != nil {
		return
	}

	// Parse username message
	var message Message
	if err := json.Unmarshal(msg, &message); err != nil {
		return
	}

	client.username = message.Content

	// Add client to connected clients
	clientsMutex.Lock()
	clients[client] = true
	clientsMutex.Unlock()

	// Broadcast user joined message
	broadcastUserUpdate("join", client.username)
	fmt.Printf("%s joined the chat!\n", client.username)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			clientsMutex.Lock()
			delete(clients, client)
			clientsMutex.Unlock()
			broadcastUserUpdate("leave", client.username)
			fmt.Printf("%s disconnected.\n", client.username)
			break
		}

		// Parse and broadcast message
		var message Message
		if err := json.Unmarshal(msg, &message); err != nil {
			continue
		}
		message.Sender = client.username
		broadcastMessage(message.Content, client)
	}
}
