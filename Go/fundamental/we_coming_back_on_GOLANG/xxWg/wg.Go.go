package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Go(func() {
			fmt.Println("worker: ", i)
		})
	}

	wg.Wait()
	fmt.Println("Done")
}

// about wg.Go, it has wg.Add() and wg.Done() by itself inside the wg.Go(func{})
