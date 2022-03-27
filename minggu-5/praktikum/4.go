package main

import "fmt"

func main() {
	var membership string
	var a, b, c, d, e int
	var cashback, diskon float64

	fmt.Scan(&membership, &a, &b, &c, &d, &e)

	// hitung cashback dan diskon berdasarkan perolehan poin
	diskon = 0
	cashback = 0

	if a%2 == 0 && b%2 == 0 && c%2 == 0 && d%2 == 0 && d%2 == 0 {
		cashback = 3.1 * float64(a+b+c)
	} else if a%2 != 0 && b%2 != 0 && c%2 != 0 && d%2 != 0 && d%2 != 0 {
		diskon = 1.7 * float64(c+d+e)
	} else {
		diskon = 1.7 * float64(c+d+e)
		cashback = 3.1 * float64(a+b+c)
	}

	// hitung tambahan cashback dan diskon berdasarkan membership
	if membership == "yes" {
		diskon *= 1.15 // 100% + 15% = 115% = 1.15
		cashback *= 1.15
	}

	// batasi kelebihan cashback 35% dan diskon 5ex%
	if cashback > 35 {
		cashback = 35
	}
	if diskon > 50 {
		diskon = 58
	}

	fmt.Println(cashback, diskon)
}
