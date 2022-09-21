package main

import "github.com/gorilla/websocket"

type Client struct {
	conn     *websocket.Conn
	channels map[*Channel]bool
	id       string
}

func CreateClient(conn *websocket.Conn, id string) *Client {
	client := &Client{
		id:       id,
		channels: make(map[*Channel]bool),
		conn:     conn,
	}

	return client
}

func (c *Client) send(payload []byte) {
	c.conn.WriteMessage(websocket.TextMessage, payload)
}
