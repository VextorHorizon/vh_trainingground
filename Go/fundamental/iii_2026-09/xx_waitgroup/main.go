package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	ch := make(chan []string)

	items := []string{"apple", "pineapple", "pencil", "mango"}

	var itemsinsomething []string
	for _, item := range items {

		itemsinsomething = append(itemsinsomething, item)

	}

	wg.Go(func() { ch <- itemsinsomething })

	fmt.Println(<-ch)
	// fmt.Println("finish everything after goroutine")
}
