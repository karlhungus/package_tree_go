package main

import (
	"fmt"
	"net"
	"bufio"
)

func main() {
	startServer(":8080")
}

func startServer(port string) {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		handleError(err)
		return
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			handleError(err)
			return
		}
		go handleConnection(conn)
	}
}

func handleError(err error) {
	fmt.Println("Error:", err)
}

func handleConnection(conn net.Conn) {
	connbuf := bufio.NewReader(conn)
	for{
		str, err := connbuf.ReadString('\n')
		if len(str)>0 {
			fmt.Println(str)
		}
		if err!= nil {
			handleError(err)
			break
		}
	}
	defer conn.Close()
}
