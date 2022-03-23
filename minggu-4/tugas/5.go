package main

import "fmt"

func checkIsPrime(n int) bool {
	var isPrime bool
	for i := 2; i <= (n / 2); i++ {
		if n%i == 0 {
			isPrime = false
		} else {
			isPrime = true
		}
	}
	return isPrime
}

func main() {
	var n, jumlahPrime int
	var isPrime bool

	fmt.Scan(&n)

	for n > 0 {

		isPrime = checkIsPrime(n)

		if isPrime {
			jumlahPrime += 1
		}

		fmt.Scan(&n)
	}

	fmt.Println("Banyak prima:", jumlahPrime)

}
