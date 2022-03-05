package main

import "fmt"

func main() {
    var r int

	fmt.Scan(&r)

	var luas_lingkaran float64 = (float64(22) / float64(7)) * float64(r)

	fmt.Println("Luas lingkaran dengan jari-jari = ", r,  "adalah ", luas_lingkaran)
}