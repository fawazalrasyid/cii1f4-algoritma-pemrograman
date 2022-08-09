package main

import "fmt"

func main() {
	fmt.Println(allOdd(9))
}

func allDigitOddRecursive(n int) bool {
	if n == 0 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	return allDigitOddRecursive(n/10)
}
