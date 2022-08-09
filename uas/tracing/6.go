func hitungGanjilGenap(n int) (int, int) {
	if n == 0 {
		return 0, 0
	}
	if n%2 == 0 {
		return hitungGanjilGenap(n/10)
	}
	return hitungGanjilGenap(n/10) + 1, hitungGanjilGenap(n/10)
}