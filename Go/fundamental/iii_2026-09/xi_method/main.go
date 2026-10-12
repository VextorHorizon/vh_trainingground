package main

import (
	"fmt"
)

type StudentName struct {
	Name  string
	Age   int
	Grade string
}

func main() {

	Veo := StudentName{Name: "Veo", Age: 24, Grade: "A+"}

	Veo.GradeSpeakUp()

}

func (s StudentName) GradeSpeakUp() {

	fmt.Printf("%s Got a grade: %s!", s.Name, s.Grade)

}
