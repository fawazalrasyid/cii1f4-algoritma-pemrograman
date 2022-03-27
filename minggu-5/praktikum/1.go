package main

import "fmt"

func main() {
	var b, faktor, nFaktor int

	fmt.Scan(&b)

	fmt.Print("Faktor: ")

	nFaktor = 0

	for faktor = 1; faktor <= b; faktor++ {
		if b%faktor == 0 {
			fmt.Print(faktor, " ")
			nFaktor += 1
		}
	}

	fmt.Println("\nPrima:", nFaktor == 2)
}
