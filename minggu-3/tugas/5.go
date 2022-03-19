package main

import "fmt"

func main() {
	var uts, uas, kuis, praktikum float64
	var nilai_akhir float64

	fmt.Scan(&uts, &uas, &kuis, &praktikum)

	if uts <= 100 && uas <= 100 && kuis <= 100 && praktikum <= 100 && uts >= 0 && uas >= 0 && kuis >= 0 && praktikum >= 0 {
		nilai_akhir = (uts * 0.35) + (uas * 0.35) + (kuis * 0.20) + (praktikum * 0.10)

		fmt.Println("Nilai akhir : ", nilai_akhir)

		if nilai_akhir > 80 {
			fmt.Println("Indeks mutu : A")
		} else if nilai_akhir > 80 {
			fmt.Println("Indeks mutu : A")
		}

	} else {
		fmt.Println("Nilai Tidak Valid")
	}

}
