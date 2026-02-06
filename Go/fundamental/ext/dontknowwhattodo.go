package main

import (
	"fmt"
)

type Task struct {
	Name   string
	isDone bool
}

func main() {

	Task1 := Task{
		Name:   "My heart",
		isDone: false,
	}
	fmt.Println("Don't know what to do today")

	if Task1.isDone == false {
		fmt.Println("Heart broken?")
	}

}
