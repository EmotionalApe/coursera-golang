package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var userName string
	fmt.Printf("Enter a name: ")
	fmt.Scan(&userName)

	var userAddress string
	fmt.Printf("Enter an address: ")
	fmt.Scan(&userAddress)

	var addressbook = make(map[string]string)
	addressbook["name"] = userName
	addressbook["address"] = userAddress

	jsonObject, err := json.Marshal(addressbook)

	if err == nil {
		fmt.Printf("Json object : %s", jsonObject)
	}

}
