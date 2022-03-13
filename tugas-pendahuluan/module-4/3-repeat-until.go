package main

import "fmt"

func main() {
	for nilaiLulus := false; !nilaiLulus; {
		fmt.Println("Kerjakan soal ujian")
		fmt.Scan(&nilaiLulus)
	}
}
