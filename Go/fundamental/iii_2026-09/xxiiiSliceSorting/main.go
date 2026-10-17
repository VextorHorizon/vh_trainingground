package main

import (
	"fmt"
	"slices"
)

func main() {

	sliceInt := []int{5, 2, 7, 9, 1, 432, 15734, 1002, 1234, 932}
	fmt.Println(sliceInt)
	slices.Sort(sliceInt)
	fmt.Println(sliceInt)

}

/*

type slice struct {
    ptr *int  // ← copy มาแล้วยังชี้โกดังเดิม
    len int
    cap int
}
*/
