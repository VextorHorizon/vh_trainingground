package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	num := rand.IntN(10)

	fmt.Println(num)

	num += num

	fmt.Println(num)
}
