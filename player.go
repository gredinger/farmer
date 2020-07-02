package farmer

import (
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
)

//Player is the user accessing the service.
type Player struct {
	gorm.Model
}

//PlayerConnection maps the Player object to the websocket connection
type PlayerConnection struct {
	User *Player
	conn *websocket.Conn
	send chan []byte
}

//writePump sends data to the server from the client
func (pc *PlayerConnection) writePump() {

}

//readsPump receives data from the client to the server
func (pc *PlayerConnection) readPump() {

}
