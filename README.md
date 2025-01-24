﻿# Go-CHAT-App
. ├── broadcasting.go ├── client.go ├── go.mod ├── go.sum ├── index.html ├── main.go └── server.go


### Files

- `broadcasting.go`: Functions for broadcasting messages and user updates.
- `client.go`: Manages client connections and handles messages.
- `index.html`: Frontend of the chat application.
- `main.go`: Entry point, sets up HTTP routes and starts the server.
- `server.go`: Manages WebSocket connections, rooms, and message broadcasting.

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/go-chat-app.git
    cd go-chat-app
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

## Usage

1. Run the server:
    ```sh
    go run main.go
    ```

2. Open your browser and navigate to `http://localhost:8080`.

## Endpoints

- `/`: Serves the [index.html](http://_vscodecontentref_/8) file.
- `/ws`: Handles WebSocket connections.
- `/active-rooms`: Returns a list of active rooms.
- `/room-users`: Returns a list of users in a specified room.

## WebSocket Messages

### Message Structure

```json
{
    "type": "message",
    "content": "Hello, World!",
    "sender": "username",
    "room": "roomName"
}

Message Types
join: User joins a room.
leave: User leaves a room.
message: User sends a message.
switch: User switches rooms.
License
This project is licensed under the MIT License.
