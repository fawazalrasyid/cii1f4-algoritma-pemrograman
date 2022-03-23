package main

import "fmt"

func main() {
	var x int
	var y bool

	fmt.Scan(&x)

	y = false

	for !y {
		if x%5 == 0 {
			y = true
			fmt.Println(y)
		} else {
			y = false
			fmt.Println(y)
			fmt.Scan(&x)
		}
	}
}
