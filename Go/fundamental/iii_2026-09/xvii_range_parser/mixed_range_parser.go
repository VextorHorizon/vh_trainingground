package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	boxSetofNumberStr := "13,10,1-12,1"

	boxSliceofNumberStr := strings.Split(boxSetofNumberStr, ",")
	fmt.Println(boxSliceofNumberStr) // [13, 10, 1-12, 1]

	fmt.Println(rangeParser(boxSliceofNumberStr))

	for _, i := range boxSliceofNumberStr {

	}

}

// 	range1 := strings.Split(rangeIwant, "-") //strings.Split() return as []string
// 	// fmt.Println(range1)

// 	firstnumStr := range1[0]
// 	endnumStr := range1[1]

// 	firstnumInt, err := strconv.Atoi(firstnumStr)
// 	if err != nil {
// 		fmt.Println(firstnumStr)
// 		fmt.Println("Error from firstnumInt Atoi")
// 	}
// 	endnumInt, err := strconv.Atoi(endnumStr)
// 	if err != nil {
// 		fmt.Println(endnumStr)
// 		fmt.Println("Error from endnumInt Atoi")
// 	}

// 	// fmt.Println(firstnumInt, endnumInt)
// 	endnumInt += 1
// 	for i := firstnumInt; i < endnumInt; i++ {
// 		fmt.Println(i)
// 	}
// }

func rangeParser(boxSlice []string) (int, int) { // can't return as int because main slice is type string (that sh)

	for _, stringthathastarget := range boxSlice {
		if strings.Contains(stringthathastarget, "-") {

			targetCut := strings.Split(stringthathastarget, "-")

			firstnumStr := targetCut[0]
			firstnumInt, err := strconv.Atoi(firstnumStr)
			if err != nil {
				fmt.Println("Error from firstnumInt in rangeParser")
			}

			endnumStr := targetCut[1]
			endnumInt, err := strconv.Atoi(endnumStr)
			if err != nil {
				fmt.Println("Error from endnumInt in rangeParser")
			}

			return firstnumInt, endnumInt
		}

	}

	return 0, 0

}

func printRangeParser(firstnum int, endnum int) {
	endnum += 1

	for i := firstnum; i < endnum; i++ {
		fmt.Println(i)
	}
}
