package main

import (
	"fmt"
	"os"
	"strings"
)

type fullName struct {
	fname string
	lname string
}

func main() {
	var namesSlice []fullName
	var fileName string
	fmt.Print("Enter the filename: ")
	fmt.Scan(&fileName)

	data, err := os.ReadFile(fileName)
	if err == nil {
		stringData := string(data)
		words := strings.Fields(stringData)

		for i := 0; i < len(words); i = i + 2 {
			nameStruct := fullName{words[i], words[i+1]}
			namesSlice = append(namesSlice, nameStruct)
		}

		for _, name := range namesSlice {
			fmt.Println(name.fname, name.lname)
		}
	}
}
