package webrtc

import (
	"encoding/json"
	"mafia/internal/core/domain"
	"mafia/internal/ports"
)

func (c *Client) WritePump() {
	for msg := range c.send {
		c.conn.WriteMessage(websocket.TextMessage, msg)
	}
}

func (c *Client) ReadPump(gameSrv ports.GameService, sfu ports.WebRTCSFU) {
	defer func() {
		sfu.RemoveClient(c)
		c.conn.Close()
	}()

	for {
		_, data, err := c.conn.ReadMessage()
		if err != nil {
			break
		}

		var msg domain.WSMessage
		if json.Unmarshal(data, &msg) != nil {
			continue
		}

		switch msg.Type {
		case "offer", "answer", "candidate":
			sfu.Broadcast(c.roomID, msg)
		case "chat":
			sfu.Broadcast(c.roomID, msg)
		}
	}
}
