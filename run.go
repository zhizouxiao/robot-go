package main

import (
	"fmt"
	"net"

	"robot-go/robot"
)

func main() {
	conn, err := net.Dial("tcp", ":12345")
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer conn.Close()
	rb := robot.NewRobot(conn)
	robot.SendLogin(rb)
	for {
		var input string
		fmt.Scanln(&input)
		// rb.Write(input + "\n")
	}

}
