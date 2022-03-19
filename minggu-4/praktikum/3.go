package main

import "fmt"

func main() {
	var n, x, jumlah int

	fmt.Scan(&n)

	for i := 1; i <= n; {
		fmt.Scan(&x)

		if x >= 0 && x < 10 {
			jumlah = jumlah + x

			i++
		}
	}

	fmt.Println(jumlah)
}
