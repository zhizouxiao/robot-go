package robot

import (
	"robot-go/robot/msg"
)

// 登录发送的第一条消息，bind_user
func SendLogin(rb *Robot) {
	_msg := msg.NewMsgRequest("bind_user", "")
	_msg.SetParam("userId", rb.userId)
	_msg.SetParam("gameId", HALL_GAMEID)
	_msg.SetParam("clientId", rb.clientId)
	rb.Write(_msg)
}

// 收到返回user_info
func ReceiveLogin(mo *msg.Msg, rb *Robot) {
	if rb.logined {
		return
	}
	cmd := mo.GetParam("cmd").(string)
	gameId := mo.GetResult("gameId").(int)
	if cmd == "user_info" {
		_msg := msg.NewMsgRequest("game", "enter")
		_msg.SetParam("userId", rb.userId)
		_msg.SetParam("gameId", HALL_GAMEID)
		_msg.SetParam("clientId", rb.clientId)
	}

	if cmd == "game_data" && gameId == HALL_GAMEID {
		_msg := msg.NewMsgRequest("game", "enter")
		_msg.SetParam("userId", rb.userId)
		_msg.SetParam("gameId", GAMEID)
		_msg.SetParam("clientId", rb.clientId)
	}

	if cmd == "game_data" && gameId == GAMEID {
		rb.gameData = mo.GetResult("result").(map[string]interface{})
		rb.logined = true
	}

}
