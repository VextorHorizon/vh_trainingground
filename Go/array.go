package main

import (
	"fmt"
)

func main() {

	var intArr [3]int
	intArr[1] = 123
	fmt.Println(intArr[2])
	fmt.Println(len(intArr))
	fmt.Println(intArr)
}
