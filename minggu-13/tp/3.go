package main

import "fmt"

// menghitung digit suatu bilangan

func digitRekursif(n int) int {

	if n%10 == n {
		return 1
	}
	return 1 + digitRekursif(n/10)
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(digitRekursif(n), "digit")
}