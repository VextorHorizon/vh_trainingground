package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	var v int
	wg.Add(3)
	for i := 1; i <= 3; i++ {

		go func() {
			defer wg.Done()
			ch <- i * 10
		}()

	}

	v += <-ch
	v += <-ch
	v += <-ch
	wg.Wait()
	fmt.Println(v)

	/*
	v = <-ch = 20
	v = <-ch = 30 (value 30 replacing 20)
	v = <-ch = 10 (value 10 replacing 30)
	*/

}
