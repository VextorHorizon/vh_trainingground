package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	userInput := args[0]

	fmt.Println(userInput)
}
