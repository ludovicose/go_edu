package main

import (
	"os"
	"fmt"
)

func main() {
	data := [][]byte{
		[]byte("The quick brown fox"),
		[]byte("jumps over the lazy dog"),
	}

	fout, err := os.Create("./filewrite.data")

	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	defer fout.Close()
	for _, out := range  data{
		fout.Write(out)
	}
}
