package main

import (
	"os"
	"fmt"
)

type metalloid struct {
	name   string
	number int32
	weight float64
}

func main() {
	var metaiiods = []metalloid{
		{"Boron", 5, 81.4},
		{"Boron1", 6, 75.5},
		{"Boron2", 7, 91.2},
	}

	file, _ := os.Create("./metalloids.txt")
	defer file.Close()

	for _, m:=range metaiiods{
		fmt.Fprint(file,"%-10s %-10d %-10.3f\n",m.name,m.number,m.weight)  // interface io.Writer
		fmt.Printf("%-10s %-10d %-10.3f\n",m.name,m.number,m.weight)  // interface stdout
	}
}
