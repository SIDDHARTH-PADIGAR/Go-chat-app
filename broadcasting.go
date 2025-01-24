package main

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

func broadcastMessage(msg string, sender *Client) {
	message := Message{
		Type:    "message",
		Content: msg,
		Sender:  sender.username,
	}

	broadcastJSON(message)
}

func broadcastUserUpdate(updateType string, username string) {
	message := Message{
		Type:    updateType,
		Content: username,
	}

	broadcastJSON(message)
}

func broadcastJSON(message Message) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()

	messageJSON, _ := json.Marshal(message)
	for client := range clients {
		client.conn.WriteMessage(websocket.TextMessage, messageJSON)
	}
}
