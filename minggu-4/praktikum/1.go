package main

import "fmt"

func main() {
	var n, a, r int

	fmt.Scan(&n, &a, &r)

	for i := 0; i <= n; i++ {
		fmt.Print(a*(i*r), " ")
	}
}
