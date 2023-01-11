package main

import (
	"log"
	"net/http"

	"websocket-simple-game-online/modules/chat"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	// ! initial map clients
	chat.Init()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.GET("/chat", chat.ChatController)

	s := http.Server{
		Addr:      ":443",
		Handler:   e,
		TLSConfig: e.Server.TLSConfig,
	}

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
