package main

import (
	"fmt"
	"os"
	"net"
	"time"
)

func main() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkErrror(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkErrror(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn)  {
	defer conn.Close()
	daytime := time.Now().String()
	conn.Write([]byte(daytime))

}

func checkErrror(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, "Fatal errror:%s", err.Error())
		os.Exit(1)
	}
}
