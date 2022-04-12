package main

import "fmt"

func negatif(a, b int, tN *int) {
	if a < 0 {
		*tN += a
	}
	if b < 0 {
		*tN += b
	}
}

func positif(a, b int, tP *int) {
	if a > 0 {
		*tP += a
	}
	if b > 0 {
		*tP += b
	}
}

func main() {
	var n, a, b, totalN, totalP int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a, &b)
		negatif(a, b, &totalN)
		positif(a, b, &totalP)
	}

	fmt.Println("Negative:", totalN)
	fmt.Println("Positive:", totalP)
}
