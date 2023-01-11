package chat

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type ChatServer struct {
	mu      sync.Mutex                 // ! สำหรับควบคุมในกรณีที่ต้องการทำอะไรสักอย่างกับ Client ใน Server ก่อน
	clients map[*websocket.Conn]string // ! ใช้ Client เป็น key และเก็บ token สำหรับป้องกันการ login ซ้อน
}

func ChatController(ctx echo.Context) error {

	// Todo |  Check Authurisezation
	token := ctx.QueryParam("token")
	if token != "token_correct" {
		return nil
	}

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	} // use default options

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(ctx.Response().Writer, ctx.Request(), nil)
	if err != nil {
		log.Print("upgrade:", err)
		return nil
	}

	defer c.Close()

	// Todo | Check User Dupicate
	chatServer.CheckUserDuplicate(token)

	chatServer.AddClient(c, token)
	defer chatServer.RemoveClient(c)

	for {
		// mt, message, err := c.ReadMessage()
		_, message, err := c.ReadMessage()

		if err != nil {
			log.Println("read:", err)
			break
		}

		log.Printf("recv: %s", message)

		// ! ส่งหา Client ที่ส่งข้อความมา
		chatServer.SendMessage(message, c)

		// ! ส่งหาคนอื่นๆ
		// chatServer.Boardcast(message)
	}
	return nil
}

func Init() {
	chatServer = &ChatServer{
		clients: make(map[*websocket.Conn]string),
	}
}

var chatServer *ChatServer
