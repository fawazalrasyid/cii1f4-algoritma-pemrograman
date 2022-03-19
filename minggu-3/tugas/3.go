package main

import "fmt"

func main() {

	var sks, bap, bap_praktikum int
	var praktikum, status_administari bool

	fmt.Scan(&sks, &bap, &praktikum, &bap_praktikum)

	if sks == 4 {
		if praktikum {
			status_administari = (float64(100)/float64(28))*float64(bap) >= 40 && (float64(100)/float64(10))*float64(bap_praktikum) >= 40
		} else {
			status_administari = (float64(100)/float64(28))*float64(bap) >= 40
		}
	} else if sks == 3 {
		if praktikum {
			status_administari = (float64(100)/float64(14))*float64(bap) >= 40 && (float64(100)/float64(10))*float64(bap_praktikum) >= 40
		} else {
			status_administari = (float64(100)/float64(14))*float64(bap) >= 40
		}
	}

	fmt.Println(status_administari)
}
