package main

import (
	"fmt"
	"os"
)

func main() {
	target, err := os.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error")
		return
	}

	fmt.Println(string(target))
}
