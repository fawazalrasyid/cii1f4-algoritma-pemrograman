package main

import "fmt"

func main() {
	var jumlah_pengunjung, jumlah_pengunjung_sebelum, jumlah_penurunan int

	fmt.Scanln(&jumlah_pengunjung)

	for jumlah_pengunjung >= 0 {

		if jumlah_pengunjung_sebelum == 0 {
			jumlah_pengunjung_sebelum = jumlah_pengunjung
		} else {
			if jumlah_pengunjung > jumlah_pengunjung_sebelum {
				jumlah_pengunjung_sebelum = jumlah_pengunjung

				fmt.Println("Status : naik")
			} else {
				jumlah_pengunjung_sebelum = jumlah_pengunjung
				jumlah_penurunan++

				fmt.Println("Status : turun")
			}
		}

		fmt.Scan(&jumlah_pengunjung)
	}

	fmt.Println("Selesai")
	fmt.Println("Terjadi penurunan", jumlah_penurunan, "kali")
}
