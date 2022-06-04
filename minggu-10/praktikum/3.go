package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func main() {
	var n, k int
	var posIdx int

	fmt.Scan(&n, &k)

	isiArray(n)

	posIdx = posisi(n, k)

	if posIdx >= 0 {
		fmt.Println(posIdx)
	} else {
		fmt.Println("TIDAK ADA")
	}

}

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	var index int

	for i := 0; i < n; i++ {
		if data[i] == k {
			index = i
		} else {
			index = -1
		}
	}

	return index
}
