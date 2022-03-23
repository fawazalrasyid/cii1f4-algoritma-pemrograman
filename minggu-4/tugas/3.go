package main

import "fmt"

func hitungDiskon(totalBelanja float64, membership bool) {
	var total_bayar float64

	if totalBelanja > 100000 && membership {
		total_bayar = totalBelanja - (totalBelanja * 0.10)
	} else if totalBelanja > 100000 && !membership {
		total_bayar = totalBelanja - (totalBelanja * 0.05)
	} else {
		total_bayar = totalBelanja
	}

	fmt.Println("Total bayar:", total_bayar)
}

func main() {
	var totalBelanja float64
	var membership bool

	fmt.Scan(&totalBelanja, &membership)

	hitungDiskon(totalBelanja, membership)

}
