package main

import "fmt"

func main() {
    var nilai int

	fmt.Scan(&nilai)

	var jumlah int = 0
	var n int = 0
	var rata2 float64 = 0

	for nilai != -1 {
		n = n + 1
		jumlah = jumlah + nilai
		fmt.Scan(&nilai)
	}

	if n == 0 {
		rata2 = 0.0
	} else {
		rata2 = float64(jumlah/n)
	}

	fmt.Println(rata2)
}