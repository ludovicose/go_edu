package main

import (
	"os"
	"fmt"
)

func main() {
	os.Mkdir("Ludovicose",0777)
	os.MkdirAll("Ludovicose/test1/test2",0777)
	err := os.Remove("Ludovicose")
	if err != nil{
		fmt.Println(err)
	}
	os.RemoveAll("Ludovicose")
}
