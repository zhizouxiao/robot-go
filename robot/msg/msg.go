package msg

import (
	"encoding/json"
)

type Msg struct {
	info map[string]interface{}
}

func (mo *Msg) GetParam(key string) interface{} {
	return mo.info[key]
}

func (mo *Msg) GetResult(key string) interface{} {
	result, ok := mo.info["result"].(map[string]interface{})
	if ok {
		return result[key]
	}
	return nil
}

func (mo *Msg) SetParam(key string, value interface{}) {
	mo.info[key] = value
}

func Pack(mo *Msg) ([]byte, error) {
	return json.Marshal(mo.info)
}

func UnPack(data []byte) (*Msg, error) {
	var msgInfo interface{}
	err := json.Unmarshal(data, &msgInfo)

	_msg := &Msg{info: msgInfo.(map[string]interface{})}
	return _msg, err
}

func NewMsgRequest(cmd string, action string) *Msg {
	msg := &Msg{info: make(map[string]interface{})}
	msg.SetParam("cmd", cmd)
	msg.SetParam("action", action)
	return msg
}
