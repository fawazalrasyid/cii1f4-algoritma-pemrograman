package main

import "fmt"

func main() {
    var s string

	fmt.Scan(&s)

	for s != "selesai" {
		fmt.Println(s)
		fmt.Scan(&s)
	}
}