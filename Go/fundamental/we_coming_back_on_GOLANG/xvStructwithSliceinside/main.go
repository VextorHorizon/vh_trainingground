package main

import (
	"fmt"
)

type StructwithSlice struct {
	Name    string
	numbers []int
}

func main() {

	Vextor := StructwithSlice{Name: "Vextor", numbers: []int{1, 2, 3, 4}}
	fmt.Println(Vextor)

	for _, number := range Vextor.numbers {
		fmt.Println(number)
	}
}
