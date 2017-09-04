package robot

import (
	. "net"
	"robot-go/robot/msg"
)

type Robot struct {
	client   *Client
	gameId   int
	userId   int
	clientId string
	gameData map[string]interface{}
	logined  bool
}

func (rb *Robot) Write(mo *msg.Msg) {
	rb.client.Write(mo)
}

func NewRobot(conn Conn) *Robot {
	return &Robot{
		gameId:   GAMEID,
		client:   newClient(conn),
		logined:  false,
		gameData: make(map[string]interface{}),
	}
}
