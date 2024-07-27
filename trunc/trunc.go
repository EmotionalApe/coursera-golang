package main

import (
	"fmt"
)

func main() {

	var userInput float32

	fmt.Printf("Please enter a floating point number: ")
	_, err := fmt.Scan(&userInput)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(int(userInput))
}
