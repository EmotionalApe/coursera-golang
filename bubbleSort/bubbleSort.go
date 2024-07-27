package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Swap(arr []int, index int) {
	var temp int = arr[index]
	arr[index] = arr[index+1]
	arr[index+1] = temp
}

func BubbleSort(arr []int) {
	var n int = len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				Swap(arr, j)
			}
		}
	}
}

func main() {
	var arr1 []int

	for i := 0; i < 10; i++ {
		var ip string
		fmt.Print("Enter a number or press X to quit: ")
		fmt.Scan(&ip)
		ip = strings.ToLower(ip)

		if ip == "x" {
			break
		} else {
			intIp, err := strconv.Atoi(ip)
			if err != nil {
				fmt.Printf("error : %x", err)
			}

			arr1 = append(arr1, intIp)
		}
	}

	BubbleSort(arr1)

	fmt.Print("The Sorted Array Is: ")
	for _, val := range arr1 {
		fmt.Print(val, " ")
	}
}
