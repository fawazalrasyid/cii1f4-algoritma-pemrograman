package main

import "fmt"

func pangkat(a, b int) int {
	var total1 int = 1

	for i := 1; i <= b; i++ {
		total1 = total1 * a
	}

	return total1
}

func konversi(biner int, desimal *int) {
	var a, j int

	for biner > 0 {
		a = biner % 10
		*desimal = *desimal + a*pangkat(2, j)
		j++
		biner = biner / 10
	}
}

func main() {
	var a, b int

	fmt.Scan(&a)

	konversi(a, &b)

	fmt.Println(b)
}
