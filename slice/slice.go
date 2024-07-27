package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sli := make([]int, 0, 3)
	for {
		var ip string
		fmt.Println("Enter X to quit or enter a number to put in the slice")
		fmt.Scan(&ip)
		if ip == "X" {
			break
		} else {
			intIp, err := strconv.Atoi(ip)
			if err != nil {
				fmt.Printf("error : %x", err)
			}

			sli = append(sli, intIp)
			sort.Ints(sli)
			fmt.Println(sli)
		}
	}
}
