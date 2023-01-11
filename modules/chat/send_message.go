package chat

import (
	"log"

	"github.com/gorilla/websocket"
)

// ! ส่งข้อความ Client
func (cs *ChatServer) SendMessage(msg []byte, c *websocket.Conn) {
	err := c.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Println("write:", err)
	}
}
