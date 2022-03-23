package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var a, b, c, d int

	fmt.Scanln(&a)
	fmt.Scanln(&b)

	rand.Seed(int64(a))

	c = rand.Intn(4) + 1
	d = rand.Intn(4) + 1

	fmt.Println(c)
	fmt.Println(d)

	if c == d {
		fmt.Println("Pemanang adalah pak Pras")
	} else if b == d {
		fmt.Println("Pemanang adalah anda")
	} else if c == d && b == d {
		fmt.Println("Seri")
	} else {
		fmt.Println("Tidak ada pemenang")
	}
}
