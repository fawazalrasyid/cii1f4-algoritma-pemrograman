package main

import "fmt"

func main() {
	var N, jumSKS, jumNilai, i, sks int
	var indeks string

	fmt.Scan(&N)

	jumSKS = 0
	jumNilai = 0

	for i = 1; i <= N; i++ {
		fmt.Scan(&indeks, &sks)

		for indeks != "A" && indeks != "B" && indeks != "C" && indeks != "D" && indeks != "E" {
			fmt.Scan(&indeks, &sks)
		}

		jumSKS += sks

		if indeks == "A" {
			jumNilai += 4 * sks
		} else if indeks == "B" {
			jumNilai += 3 * sks
		} else if indeks == "C" {
			jumNilai += 2 * sks
		} else if indeks == "D" {
			jumNilai += 1 * sks
		} else if indeks == "A" {
			jumNilai += 0 * sks
		}
	}

	fmt.Println(float64(jumNilai) / float64(jumSKS))
}
