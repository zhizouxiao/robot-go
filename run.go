package main

import (
	"fmt"
	"net"

	"robot-go/robot"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.1.254:8041")
	// conn, err := net.Dial("tcp", "192.168.1.10:8041")
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer conn.Close()
	rb := robot.NewRobot(conn)
	for {
		var input string
		fmt.Scanln(&input)
		// fmt.Println(input)
		if input == "l" {
			rb.CmdList()
		}
		// rb.Write(input + "\n")
	}

}
