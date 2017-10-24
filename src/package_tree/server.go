package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
)

var packager = NewPackager()

func main() {
	startServer(":8080")
}

func startServer(port string) {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		handleError("Listen", err)
		os.Exit(1)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			handleError("Accept", err)
			os.Exit(1)
		}
		go handleConnection(conn)
	}
}

func handleError(s string, err error) {
	fmt.Println("ERROR:", s, err)
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for{
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			handleError("Read", err)
			break;
		}
		msg := NewMessage(scanner.Text())
		conn.Write([]byte(packager.Process(msg)))
	}
}
