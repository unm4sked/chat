package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Sender  *Client
	Payload *Payload
}

func CreateMessage(client *Client, paylod *Payload) *Message {
	message := &Message{
		Sender:  client,
		Payload: paylod,
	}

	return message
}

type Payload struct {
	Text    string `json:"Text"`
	Action  string `json:"action"`
	Channel string `json:"channel"`
}

// func (payload *Payload) toJson() []byte {
// 	json, err := json.Marshal(payload)
// 	if err != nil {
// 		log.Println("Error while encode Payload", err)
// 	}

// 	return json
// }

func ParsePayload(jsonPayload []byte) (*Payload, error) {
	var payload Payload
	err := json.Unmarshal(jsonPayload, &payload)
	if err != nil {
		fmt.Println("Error while decoding text", err)
		return nil, err
	}

	return &payload, nil
}
