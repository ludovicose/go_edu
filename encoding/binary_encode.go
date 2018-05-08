package main

import (
	"time"
	"os"
	"fmt"
	"encoding/gob"
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
	books := []Book{
		Book{
			Title:"Leaning GO",
			PageCount: 375,
			ISBN: "978178439543800",
			Authors: []Name{{"Vladimir","Vivien"}},
			Publisher: "Packt",
			PublishDate: time.Date(
				2016, time.July,
				0,0,0,0,0,time.UTC,
			),
		},
		Book{
			Title:"The Go Programming Language",
			PageCount: 380,
			ISBN: "978178439543800",
			Authors: []Name{{"Alan","Donavan"},{"Brian","Kernighan"}},
			Publisher: "Packt",
			PublishDate: time.Date(
				2015, time.October,
				26,0,0,0,0,time.UTC,
			),
		},
	}
	file, err := os.Create("book.dat")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer file.Close()
	enc :=gob.NewEncoder(file)
	if err := enc.Encode(books); err != nil{
		fmt.Println(err)
	}
}
