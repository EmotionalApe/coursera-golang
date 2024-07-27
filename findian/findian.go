package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter a string: ")

	reader := bufio.NewReader(os.Stdin)
	x, _ := reader.ReadString('\n')

	x = strings.TrimSpace(x)

	x = strings.ToLower(x)
	var n int = len(x)

	if x[0] == 'i' && x[n-1] == 'n' && strings.Contains(x, "a") {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
	}
}
