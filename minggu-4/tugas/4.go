package main

import "fmt"

func hitungJumlahFibo(n int) int {
	var a, b int = 0, 1
	var jumlah int

	if n <= 1 {
		jumlah = a
	} else {
		for i := 1; i < n; i++ {
			a, b = b, a+b
			jumlah = jumlah + a
		}
	}

	return jumlah
}

func main() {
	var n, jumlahFibo int

	fmt.Scan(&n)

	jumlahFibo = hitungJumlahFibo(n)

	fmt.Println("Jumlah fibo:", jumlahFibo)

}
