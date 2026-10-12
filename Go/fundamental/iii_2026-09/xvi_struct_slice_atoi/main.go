package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ManwithNumber struct {
	name             string
	numberWithString []string
}

func main() {

	man1 := ManwithNumber{name: "manman1", numberWithString: []string{"1,2,3,4,5", "7,6,4,7,4"}}
	fmt.Println(man1) //{manman1 [1,2,3,4,5 7,6,4,7,4]}

	for _, eachone := range man1.numberWithString {

		numberString := strings.Split(eachone, ",")
		fmt.Println(numberString) // [1 2 3 4 5] a single number(string) that pack inside the box

		for _, eachnum := range numberString { //seperate number(string) one by one from the slice box
			fmt.Println(eachnum)                 // 1 2 3 4 5 in string type
			intNum, err := strconv.Atoi(eachnum) // conver number(string) to integer one by one

			fmt.Println("Real int: ", intNum) // real integer

			if err != nil {
				fmt.Println("Error from string converter eachnum")
			}
		}
	}

}
