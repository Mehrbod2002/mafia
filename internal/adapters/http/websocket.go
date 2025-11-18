package http

import (
	"mafia/internal/ports"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func WebSocketHandler(gameSrv ports.GameService, sfu ports.WebRTCSFU) gin.HandlerFunc {
	return func(c *gin.Context) {
		roomID := c.Param("room_id")
		userID := c.GetString("user_id")

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}

		client := sfu.NewClient(conn, userID, roomID)
		sfu.AddClient(roomID, client)

		go client.WritePump()
		client.ReadPump(gameSrv, sfu)
	}
}
