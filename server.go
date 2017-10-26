package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var packager = NewPackager()
var pool = NewWorkerPool(20)

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

		w := Work{
			input: conn,
			workable: func(c InputData) Result {
				handleConnection(c.(net.Conn))
				return 0
			},
		}

		pool.input <- &w

		//go handleConnection(conn)
	}
}

func handleError(s string, err error) {
	fmt.Println("ERROR:", s, err)
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for {
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			handleError("Read", err)
			break
		}
		msg := NewMessage(scanner.Text())
		fmt.Println(msg)
		result := packager.Process(msg)
		fmt.Print(result)
		_, err := conn.Write([]byte(result))
		if err != nil {
			handleError("Write", err)
			break
		}
	}
}
