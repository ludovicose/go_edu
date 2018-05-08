package main

import (
	"bytes"
	"os"
	"fmt"
)

func main() {
	var books bytes.Buffer
	books.WriteString("The Great Gatsby")
	books.WriteString("1984")
	books.WriteString("A Tale of Two  Cities")
	books.WriteString("Les Miserables")
	books.WriteString("The Call of the Wild")
	//books.WriteTo(os.Stdout)

	file, err := os.Create("./books.txt")
	if err != nil{
		fmt.Println("Unable to create file:",err)
		return
	}
	defer file.Close()
	books.WriteTo(file)

}
