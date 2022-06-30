package main

import "fmt"

func power(a, b int) int {
	if b == 0 {
		return 1
	}
	return a * power(a, b-1)
}

func main() {
	var X, Y int

	fmt.Scan(&X, &Y)

	fmt.Println(power(X, Y))
}
