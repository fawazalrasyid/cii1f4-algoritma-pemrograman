package main

import "fmt"

func fpb(a, b int) int {
	if b == 0 {
		return a
	} else {
		return fpb(b, a%b)
	}
}

func kpk(a, b int) int {
	return a / fpb(a, b) * b
}

func main() {
	var M, N int

	fmt.Scan(&M, &N)

	fmt.Println(kpk(M, N))
}
