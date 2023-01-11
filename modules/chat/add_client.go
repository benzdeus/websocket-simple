package chat

import "github.com/gorilla/websocket"

// ! เพิ่ม client
func (cs *ChatServer) AddClient(conn *websocket.Conn, token string) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	cs.clients[conn] = token
}
