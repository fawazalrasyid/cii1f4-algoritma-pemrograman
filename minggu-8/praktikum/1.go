package main

import "fmt"

func main() {
	var dadu1, dadu2, jmlGanjil int

	fmt.Scan(&dadu1, &dadu2)

	for dadu1%2 != 0 || dadu2%2 != 0 {

		if dadu1%2 != 0 && dadu2%2 != 0 {
			jmlGanjil += 1
		}

		fmt.Scan(&dadu1, &dadu2)
	}

	fmt.Println(jmlGanjil)
}
