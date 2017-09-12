package msg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Msg struct {
	info map[string]interface{}
}

func (mo *Msg) GetInfo() interface{} {
	return mo.info
}

func (mo *Msg) GetParam(key string) interface{} {
	result, ok := mo.info["params"].(map[string]interface{})
	if ok {
		return result[key]
	}
	return nil
}

func (mo *Msg) GetCmd() interface{} {
	return mo.info["cmd"]
}

func (mo *Msg) GetResult(key string) interface{} {
	result, ok := mo.info["result"].(map[string]interface{})
	if ok {
		return result[key]
	}
	return nil
}
func (mo *Msg) SetCmd(value interface{}) {
	mo.info["cmd"] = value
}

func (mo *Msg) SetParam(key string, value interface{}) {
	result, ok := mo.info["params"].(map[string]interface{})
	if ok {
		result[key] = value
	}
}

func Pack(mo *Msg) ([]byte, error) {
	return json.Marshal(mo.info)
}

func UnPack(data []byte) *Msg {
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	var x interface{}
	if err := d.Decode(&x); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("decoded to %#v\n", x)

	_msg := &Msg{info: x.(map[string]interface{})}
	return _msg
}

func NewMsgRequest(cmd string, action string) *Msg {
	msg := &Msg{info: make(map[string]interface{})}
	msg.SetCmd(cmd)
	msg.info["params"] = make(map[string]interface{})

	msg.SetParam("action", action)
	return msg
}
