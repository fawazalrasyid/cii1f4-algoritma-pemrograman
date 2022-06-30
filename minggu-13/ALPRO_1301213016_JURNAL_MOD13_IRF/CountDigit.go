package main

import "fmt"

func countDigit(n int) int {
	if n == 0 {
		return 0
	} else {
		return 1 + countDigit(n/10)
	}
}

func main() {
	var X int

	fmt.Scan(&X)

	fmt.Println(countDigit(X))
}
