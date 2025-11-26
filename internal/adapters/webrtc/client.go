package webrtc

import "mafia/internal/ports"

func (c *Client) WritePump() {}

func (c *Client) ReadPump(_ ports.GameService, _ ports.SFU) {}
