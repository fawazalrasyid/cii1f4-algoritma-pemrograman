package main

import "fmt"

// menampilkan isi elemen array

const NMAX = 10

type tabInt [NMAX]int

func printArray(arr tabInt, i int) {
	if i >= NMAX {
		return
	}
	fmt.Println(arr[i])
	printArray(arr, i+1)
}

func main() {
	var arr tabInt = [NMAX]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	printArray(arr, 0)
}
