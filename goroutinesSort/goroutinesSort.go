package main

import (
	"fmt"
	"sync"
)

func BubbleSort(arr []int, wg *sync.WaitGroup) {
	defer wg.Done()

	var n int = len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				var temp int = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
}

func merge(arr1, arr2 []int) []int {
	result := make([]int, 0, len(arr1)+len(arr2))

	i, j := 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}

	result = append(result, arr1[i:]...)
	result = append(result, arr2[j:]...)

	return result
}

func main() {
	var arr = make([]int, 0)
	fmt.Print("Enter n, number of integers you want to sort: ")
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Enter number %d: ", i+1)
		var a int
		fmt.Scan(&a)
		arr = append(arr, a)
	}

	var splitNum int = n / 4

	var splitArr1 = arr[0:splitNum]
	var splitArr2 = arr[splitNum:(2 * splitNum)]
	var splitArr3 = arr[(2 * splitNum):(3 * splitNum)]
	var splitArr4 = arr[(3 * splitNum):]

	fmt.Println()
	fmt.Printf("the 1st split array is : %v \n", splitArr1)
	fmt.Printf("the 2nd split array is : %v \n", splitArr2)
	fmt.Printf("the 3rd split array is : %v \n", splitArr3)
	fmt.Printf("the 4th split array is : %v \n", splitArr4)

	var wg sync.WaitGroup
	wg.Add(4)

	go BubbleSort(splitArr1, &wg)
	go BubbleSort(splitArr2, &wg)
	go BubbleSort(splitArr3, &wg)
	go BubbleSort(splitArr4, &wg)

	wg.Wait()

	merged1 := merge(splitArr1, splitArr2)
	merged2 := merge(splitArr3, splitArr4)

	finalMerged := merge(merged1, merged2)

	fmt.Println()
	fmt.Printf("The final array is : %v \n", finalMerged)

}
