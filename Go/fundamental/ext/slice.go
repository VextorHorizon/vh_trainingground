package main

import (
	"fmt"
)

func main() {
	fruits := []string{"Apple", "Banana"}
	fruits = append(fruits, "Pineapple")
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

}
