package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"flag"
  "log"
  "runtime/pprof"
  "runtime"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

var packager = NewPackager()

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}


	startServer(":8080")

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}
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
