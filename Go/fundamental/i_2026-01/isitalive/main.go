package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
)

func request(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if response.StatusCode >= 200 && response.StatusCode <= 300{
		fmt.Println("Website is alive!")
	} else {
		fmt.Println("Website is down?")
	}
	// fmt.Printf(response.StatusCode)

}

func main(){

	args := os.Args
	userInput := args[1]

	if userInput == "" {
		fmt.Println("Empty")
	}
	request(userInput)
	
	

}
