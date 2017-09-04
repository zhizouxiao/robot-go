package msg

import (
	"testing"
)

func TestNewMsgRequset(t *testing.T) {
	msg := NewMsgRequest("table", "info")

	if msg.GetParam("cmd") != "table" {
		t.Error("cmd error!")
	}
	if msg.GetParam("action") != "info" {
		t.Error("action error!")
	}
	t.Log(msg.GetParam("cmd").(string), msg.GetParam("action").(string))

}

func TestPackUnPack(t *testing.T) {
	msg := NewMsgRequest("table", "info")
	data, _ := Pack(msg)
	_msg, err := UnPack(data)

	if err != nil {
		t.Error("unpack error!")
	}
	t.Log(_msg.info, err, data)

}
