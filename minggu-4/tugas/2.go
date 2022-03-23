package main

import "fmt"

func checkIsTriangle(s1, s2, s3 int) bool {
	var isTriangle bool = s1+s2 > s3 && s2+s3 > s1 && s3+s1 > s2

	return isTriangle
}

func main() {
	var s1, s2, s3 int
	var isTriangle bool

	fmt.Scan(&s1, &s2, &s3)

	isTriangle = checkIsTriangle(s1, s2, s3)

	if isTriangle {
		fmt.Println("Segitiga")
	} else {
		fmt.Println("Bukan Segitiga")
	}

}
