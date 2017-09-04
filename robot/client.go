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
	client.writer.Write(data)
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
