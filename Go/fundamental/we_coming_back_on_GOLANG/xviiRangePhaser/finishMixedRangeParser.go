package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	boxSetofNumberStr := "13,10,1-12,1"

	boxSliceofNumberStr := strings.Split(boxSetofNumberStr, ",")
	// fmt.Println(boxSliceofNumberStr) // [13, 10, 1-12, 1]

	// fmt.Println(rangeParser(boxSliceofNumberStr))

	boxSetofInt := []int{}

	for _, i := range boxSliceofNumberStr {
		if strings.Contains(i, "-") {
			continue
		}
		intNum, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println("Error from intNum strconv.Atoi")
		}
		boxSetofInt = append(boxSetofInt, intNum)
	}

	firstnumInt, endnumInt := rangeParser(boxSliceofNumberStr)
	endnumInt += 1
	for i := firstnumInt; i < endnumInt; i++ {
		fmt.Printf("from rangeParser: %d\n", i)

	}

	for _, i := range boxSetofInt {
		fmt.Printf("from boxSet: %d\n", i)
	}

}

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
