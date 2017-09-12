package robot

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	. "net"
	"robot-go/robot/msg"
	"strconv"
)

type Client struct {
	reader   *bufio.Reader
	writer   *bufio.Writer
	incoming chan []byte
	seed     int
}

func (client *Client) Read() {
	buffer := []byte{}
	for {
		data := make([]byte, 65536)
		n, err := client.reader.Read(data)
		if err == io.EOF {
			continue
		}
		if client.seed == 0 {
			seed, err := strconv.ParseInt("0x"+string(data[:4]), 0, 64)
			if err != nil {
				fmt.Println("Error!", err)
			}
			client.seed = int(seed)
			// client.seed, _ = strconv.Atoi(string(data))
			mo := msg.NewMsgRequest("_connect_", "")
			msgPack, _ := msg.Pack(mo)
			client.incoming <- msgPack
			fmt.Println("read client.seed:", client.seed, seed, "0x"+string(data), err)
			continue
		}
		buffer = append(buffer, data[:n]...)
		dlen := len(buffer)
		fmt.Println("READ dlen, bufferN", dlen, n)
		if dlen > 4 {
			mlen64, err := strconv.ParseInt("0x"+string(buffer[:4]), 0, 64)
			mlen := int(mlen64)
			if err != nil {
				fmt.Println("Error!", err)
			}
			if dlen < mlen+4 {
				continue
			}
			fmt.Println("READ dlen, mlen", dlen, mlen)
			line := buffer[:mlen+4]
			buffer = buffer[mlen+4:]
			deout := decode(client.seed, line)

			var in bytes.Buffer
			in.Write(deout)
			var out bytes.Buffer
			r, _ := zlib.NewReader(&in)
			io.Copy(&out, r)
			fmt.Println("rece==============", out.String())

			client.incoming <- out.Bytes()
		}
	}

}

func (client *Client) Write(mo *msg.Msg) {
	fmt.Println("seed", client.seed)
	data, _ := msg.Pack(mo)
	endata := encode(client.seed, data)
	client.writer.Write(endata)
	client.writer.Flush()
}

func newClient(conn Conn) *Client {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	client := &Client{
		reader:   reader,
		writer:   writer,
		incoming: make(chan []byte),
		seed:     0,
	}
	go client.Read()
	return client
}

func encode(seed int, data []byte) []byte {
	mlen := []byte(fmt.Sprintf("%04x", len(data)))
	// fmt.Println(ftcode(seed+len(data), data), mlen)
	return append(mlen, ftcode(seed+len(data), data)...)
}

func decode(seed int, data []byte) []byte {
	fmt.Println("decode", len(data[4:]), seed)
	return ftcode(seed+len(data[4:]), data[4:])
}

func ftcode(seed int, data []byte) []byte {
	var randint uint = 0
	randint = uint(seed)
	out := make([]byte, len(data))

	for i := 0; i < len(data); i++ {
		randint = randint*1103515245 + 12345
		// fmt.Println(randint)
		randint = (randint / 65536) % 32768
		randchar := randint % 255
		// fmt.Println(randchar, randint)
		out[i] = byte(uint(data[i]) ^ randchar)
	}
	return out

}
