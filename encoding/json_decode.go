package main

import (
	"os"
	"fmt"
	"encoding/json"
	"time"
)

type Name struct {
	First,Last string
}

type Book struct {
	Title string
	PageCount int
	ISBN string
	Authors []Name
	Publisher string
	PublishDate time.Time
}


func main() {

	file, err := os.Open("book.dat")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer file.Close()

	var books []Book
	dec := json.NewDecoder(file)
	if err := dec.Decode(&books); err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println(books)
}
