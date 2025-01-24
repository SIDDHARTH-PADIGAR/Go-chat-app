// main.go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/ws", handleConnection)
	http.HandleFunc("/active-rooms", handleActiveRooms)
	http.HandleFunc("/room-users", handleRoomUsers)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
