package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Empty")
		return
	}

	fmt.Printf("User input is: %v", args[1])

}
