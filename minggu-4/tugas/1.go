package main

import "fmt"

func konversiWaktu(jam, menit, detik int) int {
	var hasilKonversi int

	hasilKonversi = (jam * 3600) + (menit * 60) + detik

	return hasilKonversi
}

func main() {
	var jam, menit, detik, hasil int

	fmt.Scan(&jam, &menit, &detik)

	hasil = konversiWaktu(jam, menit, detik)

	fmt.Println("Hasil konversi:", hasil)
}
