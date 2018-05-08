package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

func main() {
	file, err := os.Open("./filewrite.data")
	if err != nil{
		fmt.Println("Unable to open file:",err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for{
		line,err := reader.ReadString('\n')
		if err != nil{
			if err == io.EOF{
				break
			} else {
				fmt.Println("Error reading",err)
				return
			}
		}
		fmt.Println(line)
	}
}
