package main

import "fmt"

func main() {
	var punyaSim bool

	fmt.Scan(&punyaSim)

	for punyaSim {
		fmt.Println("Boleh mengemudi")
		fmt.Scan(&punyaSim)
	}
}
