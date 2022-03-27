package main

import "fmt"

func main() {
	var gram, kg, gr, hargaKg, hargaGr, total int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&gram)

	kg = gram / 1000
	gr = gram % 1000
	fmt.Println("Detail berat:", kg, "kg +", gr, "gr")

	hargaKg = 10000 * kg
	if gr > 500 {
		hargaGr = gr * 5
	} else {
		hargaGr = gr * 15
	}
	fmt.Println("Detail biaya: Rp.", hargaKg, "+ Rp.", hargaGr)

	if kg > 10 {
		total = hargaKg
	} else {
		total = hargaKg + hargaGr
	}
	fmt.Println("Total biaya: Rp.", total)
}
