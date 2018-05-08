package main

import (
	"os"
	"fmt"
	"io"
)

func main() {
	fin, err := os.Open("./filewrite.data")
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer fin.Close()
	p := make([]byte,1024)
	for{
		n,err := fin.Read(p)
		if err == io.EOF{
			break
		}
		fmt.Println(string(p[:n]))
	}
}
