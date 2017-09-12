package msg

import (
	"testing"
)

func TestNewMsgRequset(t *testing.T) {
	msg := NewMsgRequest("table", "info")
	t.Log(msg.GetInfo())

	if msg.GetCmd() != "table" {
		t.Error("cmd error!")
	}
	if msg.GetParam("action") != "info" {
		t.Error("action error!")
	}
	t.Log(msg.GetCmd().(string), msg.GetParam("action").(string))

}

func TestPackUnPack(t *testing.T) {
	msg := NewMsgRequest("table", "info")
	data, _ := Pack(msg)
	_msg := UnPack(data)

	t.Log(_msg.info, data)

}

func TestPackUnPack2(t *testing.T) {
	data := `{"cmd":"user_info","result":{"gameId":9999,"userId":101,"udata":{"email":"","pdevid":"","mdevid":"","isbind":0,"snsId":"","name":"\u7687\u752b\u6210\u534e","source":"","diamond":0,"address":"","sex":0,"state":0,"payCount":0,"snsinfo":"","vip":0,"dayang":0,"idcardno":"","phonenumber":"","truename":"","detect_phonenumber":"","lang":"","country":"","signature":"","set_name_sum":0,"coupon":0,"purl":"http://ddz.image.tuyoo.com/avatar/head_horse.png","beauty":0,"charm":0,"password":"","bindMobile":"","pt":0,"createTime":"2017-08-08 16:32:24.568555","coin":0,"chip":1115,"isBeauty":false,"vipInfo":{"level":0,"name":"VIP0","exp":0,"expCurrent":0,"expNext":60},"assistance":{"count":0,"limit":1000}},"loc":"0.0.0.0"}}`
	_msg := UnPack([]byte(data))

	t.Log(_msg.info, data)

}
