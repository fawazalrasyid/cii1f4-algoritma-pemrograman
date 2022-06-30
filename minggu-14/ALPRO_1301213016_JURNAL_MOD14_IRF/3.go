package main

import "fmt"

const mark = '.'
const space = '-'

var pita string
var CC byte
var EOP bool
var indeks int

var LKata int
var LTotal int
var NKata int

func main() {
	fmt.Scan(&pita)
	LTotal = 0
	NKata = 0
	startKata()

	for LKata != 0 {
		LTotal += LKata
		NKata++
		advKata()
	}

	if NKata != 0 {
		fmt.Printf("%.2f", float64(LTotal)/float64(NKata))

	} else {
		fmt.Println("Pita Kosong")
	}
}

func START() {
	indeks = 0
	CC = pita[indeks]
	EOP = CC == mark
}

func ADV() {
	indeks++
	CC = pita[indeks]
	EOP = CC == mark
}

func ignoreBlank() {
	for CC == space {
		ADV()
	}
}

func hitungPanjang() {
	LKata = 0
	for CC != space && !EOP {
		LKata++
		ADV()
	}
}

func startKata() {
	START()
	ignoreBlank()
	hitungPanjang()
}

func advKata() {
	ignoreBlank()
	hitungPanjang()
}
