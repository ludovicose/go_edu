package main

import "fmt"

func main() {
	var choice int
	fmt.Println("A square is what?")
	fmt.Print("Enter 1=quarilatereal 2 rectagonal:")

	n, err := fmt.Scanf("%d",&choice)
	if n != 1 || err != nil{
		fmt.Println("follow directions!")
		return
	}

	if choice == 1 {
		fmt.Println("You are correct!")
	} else {
		fmt.Println("Wrong, Google it")
	}
}
