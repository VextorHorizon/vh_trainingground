package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	slice := []int{1, 2, 5, 7, 9, 14}
	fmt.Println(slice)

	for i, number := range slice {
		i += 1
		fmt.Printf("%d) %d \n", i, number)
	}

	amount_slice := len(slice)

	fmt.Printf("Amount of integer in slice is %v", amount_slice)
}
