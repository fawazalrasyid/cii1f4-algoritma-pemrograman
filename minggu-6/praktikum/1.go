package main

import "fmt"

func findFactorial(n int, result *int) {
	var j int
	*result = 1

	for j = 1; j <= n; j++ {
		*result *= j
	}
}

func permutation(n, r int) int {
	var nF, nrF int
	findFactorial(n, &nF)
	findFactorial(n-r, &nrF)

	return nF / nrF
}

func combination(n, r int) int {
	var nF, nrF, rF int
	findFactorial(n, &nF)
	findFactorial(n-r, &nrF)
	findFactorial(n, &nF)
	findFactorial(r, &rF)

	return nF / (rF * nrF)
}

func main() {
	var a, b, c, d int
	var p1, c1, p2, c2 int

	fmt.Scan(&a, &b, &c, &d)
	p1 = permutation(a, c)
	c1 = combination(a, c)
	p2 = permutation(b, d)
	c2 = combination(b, d)

	fmt.Println(p1, c1)
	fmt.Println(p2, c2)
}
