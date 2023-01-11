package chat

import "github.com/gorilla/websocket"

// ! ลบ client
func (cs *ChatServer) RemoveClient(conn *websocket.Conn) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	delete(cs.clients, conn)
}
