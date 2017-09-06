package robot

import (
	"bufio"
	"fmt"
	"io"
	. "net"
	"robot-go/robot/msg"
)

type Client struct {
	reader *bufio.Reader
	writer *bufio.Writer
	seed   int
}

func (client *Client) Read() {
	for {
		line, err := client.reader.ReadString('\n')
		if err == io.EOF {
			continue
		}
		fmt.Println("read:", line)
	}

}

func (client *Client) Write(mo *msg.Msg) {
	data, _ := msg.Pack(mo)
	endata := ftcode(client.seed, data)
	client.writer.Write(endata)
	client.writer.Flush()
}

func newClient(conn Conn) *Client {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	client := &Client{
		reader: reader,
		writer: writer,
	}
	go client.Read()
	return client
}

func ftcode(seed int, data []byte) []byte {
	var randint uint = 0
	randint = uint(seed)
	out := make([]byte, len(data))

	for i := 0; i < len(data); i++ {
		randint = randint*1103515245 + 12345
		fmt.Println(randint)
		randint = (randint / 65536) % 32768
		randchar := randint % 255
		fmt.Println(randchar, randint)
		out[i] = byte(uint(data[i]) ^ randchar)
	}
	return out

}
