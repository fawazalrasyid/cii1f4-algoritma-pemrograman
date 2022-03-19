package main

import "fmt"

func main() {
	var N, a, b int
	var x int

	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		fmt.Scan(&a, &b)
		x = 1
		u := 1
		for u <= b {
			x = x * a
			u += 1
		}
		fmt.Println(x)
	}
}
