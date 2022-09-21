package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const (
	sendMessage   = "send-message"
	createChannel = "create-channel"
	joinChannel   = "join-channel"
	listChannels  = "list-channels"
	quitChannel   = "quit-channel"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type WSServer struct {
	channels map[*Channel]bool
	clients  map[*Client]bool
}

func CrateWebSocketServer() *WSServer {
	server := &WSServer{
		clients:  make(map[*Client]bool),
		channels: make(map[*Channel]bool),
	}

	return server
}

func (ws *WSServer) AddClient(client *Client) {
	ws.clients[client] = true
}

func FireWsServer(server *WSServer, w http.ResponseWriter, r *http.Request) {

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != err {
		log.Println(err)
	}

	id := uuid.New().String()
	log.Println("Client successfully Connected... with id: ", id)

	client := CreateClient(conn, id)
	server.AddClient(client)

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		payload, err := ParsePayload(p)
		if err != nil {
			fmt.Println(err)
		}

		message := CreateMessage(client, payload)

		server.handleMessage(message)

		fmt.Println("Message from client: ", string(p))
	}
}

// {text: xxx, action: "action1"}
func (w *WSServer) handleMessage(message *Message) {
	fmt.Println("Action: ", message.Payload.Action)
	switch message.Payload.Action {
	case sendMessage:
		{
			for channel := range message.Sender.channels {
				channel.sendMessage([]byte(message.Payload.Text))
			}
		}
	case createChannel:
		{
			newChannel := CreateChannel(message.Payload.Channel)
			w.channels[newChannel] = true
		}
	case joinChannel:
		{
			var channel *Channel
			for ch := range w.channels {
				if ch.name == message.Payload.Channel {
					channel = ch
				}
			}

			message.Sender.channels[channel] = true
		}
	case listChannels:
		{
			var channels []string
			for channel := range w.channels {
				channels = append(channels, channel.name)
			}
			fmt.Println("List of channels", channels)
			jsonResult, err := json.Marshal(channels)
			if err != nil {
				fmt.Println("Stringify Error", err)
				break
			}

			err = message.Sender.conn.WriteMessage(websocket.TextMessage, jsonResult)
			if err != nil {
				fmt.Println("Write message error", err)
			}
		}
	case quitChannel:
		{
			var channel *Channel
			for ch := range w.channels {
				if ch.name == message.Payload.Channel {
					channel = ch
				}
			}
			if channel == nil {
				break
			}

			delete(message.Sender.channels, channel)
		}

	default:
		fmt.Println("Action not match: ", message.Payload.Action)
	}
}
