package main

import (
	"fmt"
	"slices"
)

type slicez struct {
	name     string
	sliceInt []int
}

func main() {
	slice1 := slicez{
		name:     "Peter",
		sliceInt: []int{5, 2, 7, 9, 1, 432, 15734, 1002, 1234, 932},
	}

	fmt.Println(slice1)
	slices.Sort(slice1.sliceInt)
	fmt.Println(slice1)

}

/*

type slice struct {
    ptr *int  // ← copy มาแล้วยังชี้โกดังเดิม
    len int
    cap int
}
*/
