package robot

import (
	"testing"
)

func TestFtcode(t *testing.T) {
	seed := 123456
	data := []byte("abcdefg")
	out := ftcode(seed+len(data), data)
	t.Log(out, string(out))

	deout := ftcode(seed+len(out), out)
	t.Log(deout, string(deout))
	t.Error("test")
}
