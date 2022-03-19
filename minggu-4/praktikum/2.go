package main

import "fmt"

func main() {

	var n, d1, d0 int
	var status bool = true

	fmt.Scan(&n)

	d1 = n % 10
	d0 = d1
	n = n % 10

	for n > 0 && status {
		d1 = n % 10
		status = (d1-d0) == 1 || (d0-d1) == 1
		d0 = d1
		n = n / 10
	}

	fmt.Print(!status)
}
