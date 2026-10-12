package main

import (
	"fmt"
)

func main() {
	var x int // 0
	for x < 10 {
		x = plusnum(x)
		fmt.Println(x)
	}

	// 	for x < 10 {
	// 		x := plusnum(x)    // +1
	// 		fmt.Println("", x) // +1
	// 	}
}

func plusnum(x int) int {
	return x + 1
}
