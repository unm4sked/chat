# Hi


 ws/channels

 
-----------------
broadcast => todo


 type Blab struct {
    Action => 
    Message => ""
    Channel => multi channels ?
 }

server {
    maps of channels
}

channel{
    name
    clients []
}

Action:
- sendMessage
- createChannel
- joinChannel
- listChannels
- quitChannel



1. Boilerplate ws  [done]
2. Create structure => Channel(Room) , Message(Payload), Client, server
3. Implement Actions for createRoom joinChannel, sendMessage
4. 