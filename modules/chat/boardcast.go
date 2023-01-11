package chat

import (
	"log"

	"github.com/gorilla/websocket"
)

// ! ส่งข้อความหาทุกคน
func (cs *ChatServer) Boardcast(msg []byte) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	for conn := range cs.clients {
		err := conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println(err)
			continue
		}

	}
}
