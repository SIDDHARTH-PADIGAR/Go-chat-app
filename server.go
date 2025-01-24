package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Room struct {
	clients map[*websocket.Conn]bool
	mutex   sync.Mutex
}

var (
	rooms      = make(map[string]*Room)
	roomsMutex sync.RWMutex
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func newRoom() *Room {
	return &Room{clients: make(map[*websocket.Conn]bool)}
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
	}
	defer conn.Close()

	var message Message
	err = conn.ReadJSON(&message)
	if err != nil {
		log.Println("Error reading message: ", err)
		return
	}

	roomName := message.Room

	roomsMutex.Lock()
	if _, exists := rooms[roomName]; !exists {
		rooms[roomName] = newRoom()
	}
	room := rooms[roomName]
	room.clients[conn] = true
	roomsMutex.Unlock()

	log.Printf("Client %s joined room: %s\n", message.Sender, roomName)

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			break
		}
		msg.Sender = message.Sender
		msg.Room = roomName
		broadcastToRoom(roomName, msg)
	}

	roomsMutex.Lock()
	delete(room.clients, conn)
	roomsMutex.Unlock()
}

func broadcastToRoom(roomName string, msg Message) {
	roomsMutex.RLock()
	room, exists := rooms[roomName]
	roomsMutex.RUnlock()

	if !exists {
		return
	}

	room.mutex.Lock()
	defer room.mutex.Unlock()

	for client := range room.clients {
		err := client.WriteJSON(msg)
		if err != nil {
			client.Close()
			delete(room.clients, client)
		}
	}
}

func handleActiveRooms(w http.ResponseWriter, r *http.Request) {
	roomsMutex.RLock()
	defer roomsMutex.RUnlock()

	activeRooms := make([]string, 0, len(rooms))
	for roomName := range rooms {
		activeRooms = append(activeRooms, roomName)
	}

	w.Header().Set("Content=Type", "application/json")
	json.NewEncoder(w).Encode(activeRooms)
}

func handleRoomUsers(w http.ResponseWriter, r *http.Request) {
	roomName := r.URL.Query().Get("room")
	if roomName == "" {
		http.Error(w, "Room not specified", http.StatusBadRequest)
	}

	roomsMutex.RLock()
	room, exists := rooms[roomName]
	roomsMutex.RUnlock()

	if !exists {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	room.mutex.Lock()
	users := make([]string, 0, len(room.clients))
	for client := range room.clients {
		//This is a placeholder. In a real application, you would get the username from the clients map.
		users = append(users, fmt.Sprintf("User %d", client))
	}
	room.mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
