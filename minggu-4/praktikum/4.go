package main

import (
	"fmt"
)

func main() {
	var n int
	var topi, alatTulis, tali, pisau bool
	var statusKesiapan = true

	fmt.Scan(&n)

	for i := 1; i <= n && statusKesiapan; i++ {
		fmt.Scan(&topi, &alatTulis, &tali, &pisau)

		statusKesiapan = topi && alatTulis && tali && pisau

	}

	fmt.Println(statusKesiapan)
}
