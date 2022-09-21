package main

type Channel struct {
	name    string
	clients map[*Client]bool
}

func CreateChannel(name string) *Channel {

	var clients = make(map[*Client]bool)
	return &Channel{
		name:    name,
		clients: clients,
	}
}

func (channel *Channel) sendMessage(message []byte) {
	for client, _ := range channel.clients {
		client.send(message)
	}
}

func (channel *Channel) removeClient(client *Client) {
	if _, ok := channel.clients[client]; ok {
		delete(channel.clients, client)
	}
}

func (channel *Channel) addClient(client *Client) {
	channel.clients[client] = true
}
