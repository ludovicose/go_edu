package main

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	rows := []string{
		"The quick brown fox",
		"jums over the lazy dog",
	}

	fout, err := os.Create("./filewrite.data")
	writer := bufio.NewWriter(fout)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	defer fout.Close()

	for _, row := range rows {
		writer.WriteString(row)
	}
	writer.Flush()
}
