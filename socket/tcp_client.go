package main

import (
	"net"
	"fmt"
	"os"
	"io/ioutil"
)

func main() {

	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkErrror(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkErrror(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkErrror(err)
	result, err := ioutil.ReadAll(conn)
	checkErrror(err)
	fmt.Println(string(result))
	os.Exit(0)

}

func checkErrror(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, "Fatal errror: %s", err.Error())
		os.Exit(1)
	}
}
