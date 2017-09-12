package robot

import (
	"fmt"
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

func (rb *Robot) CmdList() {
	fmt.Println("1. quick_start")
	fmt.Println("2. table_list")
}

func NewRobot(conn Conn) *Robot {
	client := newClient(conn)

	rb := &Robot{
		gameId:   GAMEID,
		userId:   101,
		client:   client,
		logined:  false,
		gameData: make(map[string]interface{}),
		clientId: "IOS_4.01_weixinPay,tyGuest,tyAccount.alipay.0-hall8.tuyoo.tu",
	}

	go func() {
		for {
			select {
			case data := <-client.incoming:
				fmt.Println("Receive", string(data))
				mo := msg.UnPack(data)
				fmt.Println("Receive", mo.GetInfo())
				ReceiveLogin(mo, rb)

			}
		}
	}()
	return rb
}
