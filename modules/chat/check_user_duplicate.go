package chat

import "log"

// ! เช็ค login ซ้อน
func (cs *ChatServer) CheckUserDuplicate(tokenUser string) {
	cs.mu.Lock()
	defer cs.mu.Unlock()

	// ! สิ่งที่ต้อง implement เพิ่มคือ ถอด Token เพื่อเช็ค user อีกที
	for conn, token := range cs.clients {
		log.Println(token)
		if token == tokenUser {
			conn.Close()
			break
		}
	}
}
