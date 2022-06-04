package main

import "fmt"

type calon struct {
	no   int
	vote int
}

func main() {
	var x int
	var calon [20]calon
	var max, vote_ketua int
	var ketua, wakil int

	k := 1
	for i := 0; i < 20; i++ {
		calon[i].no = k
		k++
	}

	total := 0
	sah := 0

	fmt.Scan(&x)
	for x != 0 {
		if x <= 20 && x > 0 {
			sah++
			calon[x-1].vote++
		}
		total++
		fmt.Scan(&x)
	}
	fmt.Println("Suara masuk:", total)
	fmt.Println("Suara sah:", sah)

	max = 0

	for i := 0; i < 20; i++ {
		if calon[i].vote > max {
			ketua = i + 1
			max = calon[i].vote
			vote_ketua = calon[i].vote
		}
	}

	for i := 0; i < 20; i++ {
		if calon[i].vote >= vote_ketua {
			wakil = i + 1

			fmt.Print(wakil)
		}
	}

	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}
