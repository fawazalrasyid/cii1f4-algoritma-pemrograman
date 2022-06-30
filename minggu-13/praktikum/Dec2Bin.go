package main

import "fmt"

func decimalToBinary(n int) string {
	if n == 0 {
		return "0"
	} else if n == 1 {
		return "1"
	} else {
		if n%2 == 1 {
			return decimalToBinary(n/2) + "1"
		} else {
			return decimalToBinary(n/2) + "0"
		}
	}
}

func main() {
	var X int

	fmt.Scan(&X)

	fmt.Println(decimalToBinary(X))
}
