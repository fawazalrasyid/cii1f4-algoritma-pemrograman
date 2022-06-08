package main

import "fmt"

// konversi bilangan desimal menjadi binary

func biner(d int) int {
	if d == 1 {
		return 1
	} else if d == 0 {
		return 0
	}

	mod := d % 2
	div := d / 2

	return mod + (biner(div) * 10)
}

func main() {
	var n int

	fmt.Scan(&n)

	fmt.Println(biner(n))
}
