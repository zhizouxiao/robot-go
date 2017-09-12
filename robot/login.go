package robot

import (
	"encoding/json"
	"fmt"
	"robot-go/robot/msg"
	"strconv"
)

// 登录发送的第一条消息，bind_user
func SendLogin(rb *Robot) {
	_msg := msg.NewMsgRequest("bind_user", "")
	_msg.SetParam("userId", rb.userId)
	_msg.SetParam("gameId", HALL_GAMEID)
	_msg.SetParam("clientId", rb.clientId)
	fmt.Println("SendLogin==", _msg.GetInfo())
	rb.Write(_msg)
}

// 收到返回user_info
func ReceiveLogin(mo *msg.Msg, rb *Robot) {
	fmt.Println("ReceiveLogin==", rb.logined)
	if rb.logined {
		return
	}

	cmd := mo.GetCmd().(string)
	if cmd == "_connect_" {
		SendLogin(rb)
		return
	}
	fmt.Println("gameId:", mo.GetResult("gameId"))
	number, ok := mo.GetResult("gameId").(json.Number)
	if !ok {
		return
	}
	gameId, _ := strconv.Atoi(string(number))
	if cmd == "user_info" {
		_msg := msg.NewMsgRequest("game", "enter")
		_msg.SetParam("userId", rb.userId)
		_msg.SetParam("gameId", HALL_GAMEID)
		_msg.SetParam("clientId", rb.clientId)
		rb.Write(_msg)
	}

	if cmd == "game_data" && gameId == HALL_GAMEID {
		_msg := msg.NewMsgRequest("game", "enter")
		_msg.SetParam("userId", rb.userId)
		_msg.SetParam("gameId", GAMEID)
		_msg.SetParam("clientId", rb.clientId)
		rb.Write(_msg)
	}

	if cmd == "game_data" && gameId == GAMEID {
		// rb.gameData = mo.GetResult("result").(map[string]interface{})
		rb.logined = true
	}

}
