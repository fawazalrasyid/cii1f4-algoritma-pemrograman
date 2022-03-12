package main

import "fmt"

func main() {

	var teks1, teks2, teks3 string
	var hasil bool

	fmt.Scan(&teks1, &teks2, &teks3)

	hasil = teks1 == teks2 || teks1 == teks3 || teks2 == teks3

	fmt.Println(hasil)
}
