package main

import "fmt"

func main() {
	var N, M, d, total int
	var f bool

	fmt.Scan(&N, &M)
	total = 0

	fmt.Scan(&d)

	f = (((M % 10000 / 1000) == d) || (M%1000/100) == d || ((M % 100 / 10) == d) || ((M % 10 / 1) == d))

	for i := 0; i < N && f; i++ {
		total = total + 1
		fmt.Scan(&d)
		f = (((M % 10000 / 1000) == d) || (M%1000/100) == d || ((M % 100 / 10) == d) || ((M % 10 / 1) == d))
	}

	fmt.Println(total)
}
