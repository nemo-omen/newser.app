package handler

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type WsHandler struct{}

func (h WsHandler) HandleWsConnect(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	type msg struct {
		data string
	}

	conn.WriteJSON(msg{data: "livereload"})
	return nil
}
