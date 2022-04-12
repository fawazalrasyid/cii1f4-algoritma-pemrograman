package main

import "fmt"

func main() {
	var nama, nama_min, nama_max string
	var p int
	var faktor, jumlah, min, mean, max float64

	nama_min = ""
	nama_max = ""

	fmt.Scanln(&nama)

	for nama != "STOP" {
		fmt.Scanln(&p)
		for i := 0; i < 3*p; i++ {
			fmt.Scanln(&faktor)
			jumlah += faktor
		}
		mean = jumlah / float64(3*p)
		if nama_min == "" || nama_min != "" && min > mean {
			min = mean
			nama_min = nama
		}
		if nama_max == "" || nama_max != "" && max < mean {
			max = mean
			nama_max = nama
		}
		fmt.Scan(&nama)
		jumlah = 0
	}

	fmt.Printf("%s %2.f", nama_max, max)
	fmt.Printf("\n%s %2.f", nama_min, min)
}
