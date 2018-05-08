package main

import (
	"os"
	"fmt"
	"runtime"
)

func main() {
	userFile := "Ludovicose.txt"
	fout, err := os.Create(userFile)
	if err != nil{
		fmt.Println(err)
		runtime.ReadTrace()
	}
	defer fout.Close()
	for i:=0; i<10; i++{
		fout.WriteString("Just a test \r\t")
		fout.Write([]byte("Just a test! \r\t"))
	}
}
