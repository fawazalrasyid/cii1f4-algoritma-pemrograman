package main

import "fmt"

func main() {
	var n, juri1A, juri2A, juri3A, juri1B, juri2B, juri3B, skorA, skorB int

	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&juri1A, &juri2A, &juri3A, &juri1B, &juri2B, &juri3B)

		skorA += juri1A + juri2A + juri3A
		skorB += juri1B + juri2B + juri3B
	}

	fmt.Println("Petinju A:", skorA, "dan petinju B:", skorB)

	if skorA > skorB {
		fmt.Println("Pemenang adalah petinju A")
	} else if skorB > skorA {
		fmt.Println("Pemenang adalah petinju B")
	} else {
		fmt.Println("Draw")
	}
}
