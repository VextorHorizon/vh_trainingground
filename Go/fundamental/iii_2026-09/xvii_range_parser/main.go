package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	rangeIwant := "1-12"
	range1 := strings.Split(rangeIwant, "-")
	// fmt.Println(range1)

	firstnumStr := range1[0]
	endnumStr := range1[1]

	firstnumInt, err := strconv.Atoi(firstnumStr)
	if err != nil {
		fmt.Println("Error from firstnumInt Atoi")
	}
	endnumInt, err := strconv.Atoi(endnumStr)
	if err != nil {
		fmt.Println("Error from endnumInt Atoi")
	}

	// fmt.Println(firstnumInt, endnumInt)
	endnumInt += 1
	for i := firstnumInt; i < endnumInt; i++ {
		fmt.Println(i)
	}
}
