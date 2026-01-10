package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	var numA int = 10
	var numB int = 12
	var result int = numA + numB

	fmt.Printf("Number is %d\n", result)
	var printhaha = "hellohello"
	printhelloworld(printhaha)

	plusnumberfuncresult := plusnumber(numA, numB)
	fmt.Printf("%d", plusnumberfuncresult)

}

func printhelloworld(printvalue string) {
	fmt.Println(printvalue)
}

func plusnumber(numbera, numberb int) int { // int ข้างหลังคือเป็นสัญญาผูกไว้ว่า return เป็น int
	result := numbera + numberb
	return result //หรือว่าจะเขียนว่า return numbera + numberb ได้เลยก้ได้นะ
}
