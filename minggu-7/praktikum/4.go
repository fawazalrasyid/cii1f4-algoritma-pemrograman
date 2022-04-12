package main

import "fmt"

func museum(jumlah *int, i *int) {
	var penurunan int
	var pengunjung, temp int
	*i = 1
	for penurunan != 3 {
		fmt.Printf("Hari %d: ", *i)
		fmt.Scanln(&pengunjung)
		for pengunjung < 0 || pengunjung > 100 {
			fmt.Printf("Hari %d: ", *i)
			fmt.Scanln(&pengunjung)
		}
		*jumlah = *jumlah + pengunjung
		cek(pengunjung, temp, &penurunan)
		temp = pengunjung
		*i++
	}
}

func cek(pengunjung, temp int, penurunan *int) {
	if pengunjung < temp {
		*penurunan++
	} else {
		*penurunan = 0
	}
}

func main() {
	var i, jumlah int

	museum(&jumlah, &i)
	
	fmt.Printf("Museum buka selama %d hari", i-1)
	fmt.Printf("\nRata-rata pengunjung %.2f orang", float64(jumlah)/float64(i-1))
}