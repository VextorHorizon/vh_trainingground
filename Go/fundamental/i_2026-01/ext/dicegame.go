package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	fmt.Println("Choose a number between 1-10: ")

	var userinput int
	fmt.Print("Your guessing number: ")
	// _, err := fmt.Scanf("%d", &userinput)
	// // _, err := fmt.Scanf("Your guessing number: %d", &userinput) ไม่ได้นะจะน้อง scanf มันจะอ่าน input ข้างในหมดเลย ก้จะได้ input เป็น "Your guessing number: ..."
	// if err != nil {
	// 	fmt.Println("Use a integer")
	// 	return
	// }
	_, err := fmt.Scanln(&userinput)
	if err != nil {
		if err.Error() == "unexpected newline" {
			fmt.Print("User input is empty")
		} else {
			fmt.Print("Use a integer")
		}
		return
	}
	if userinput > 10 || userinput < 2 {
		fmt.Print("Number limit is only 1-10")
		return
	}

	randomnumber(userinput)

}

func randomnumber(userint int) {
	//function ให้ user ทายเลข
	var randomnum int = rand.IntN(10) + 1 // +1 เพราะว่าให้ range มันขยับจาก 0-9 ไป 1-10
	if userint == randomnum {
		fmt.Printf("Congratulation! You guess it! Lucky number is: %d", userint)
	} else {
		fmt.Printf("Still miss the guess huh? the lucky number is %d", randomnum)
	}
}

// first time on writing Go basic project

