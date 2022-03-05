package main

import "fmt"

func main() {
    var r int

	fmt.Scan(&r)

	luas_lingkaran := 22/7 * float64(r)

	fmt.Println("Luas lingkaran dengan jari-jari = ", r,  "adalah ", luas_lingkaran)
}