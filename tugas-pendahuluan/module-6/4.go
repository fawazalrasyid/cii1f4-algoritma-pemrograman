package main

import "fmt"

func main() {
	// Subprogram mencari nilai maksimum dan minimum
	var x1, x2, x3, x4 int = 3, 5, 2, 9

	// 4.a function
	fmt.Println("Pemanggilan function")
	// variable berguna untuk menampung nilai kembalian dari function (function ini memeliki 2 nilai kembalian)
	var min, max int = findMinMaxFunc(x1, x2, x3, x4)
	// print hasil
	fmt.Println("Nilai minimum:", min, "Nilai maximum:", max)

	// 4.a procedure
	fmt.Println("\nPemanggilan procedure")
	// pemanggilan procedure (tidak perlu ditampung dalam variable karena procedure tidak memiliki nilai kembalian)
	findMinMaxPro(x1, x2, x3, x4)

	// ___________________________________________________________________________________________________________

	// Subprogram bilangan bulat ke biner
	var x int = 1032

	// 4.b function
	fmt.Println("\nPemanggilan function")
	// variable berguna untuk menampung nilai kembalin dari function
	var binary string = convertToBinaryFunc(x)
	// print hasil
	fmt.Println(binary)

	// 4.b procedure
	fmt.Println("\nPemanggilan procedure")
	// pemanggilan procedure (tidak perlu ditampung dalam variable karena procedure tidak memiliki nilai kembalian)
	convertToBinaryPro(x)

}

func findMinMaxFunc(x1, x2, x3, x4 int) (int, int) {
	var arr [4]int = [4]int{x1, x2, x3, x4}

	var min int = x1
	var max int = x1

	for i := 1; i < 4; i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}

	return min, max

}

func findMinMaxPro(x1, x2, x3, x4 int) {
	var arr [4]int = [4]int{x1, x2, x3, x4}

	var min int = x1
	var max int = x1

	for i := 1; i < 4; i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}

	fmt.Println("Nilai minimum:", min, "Nilai maximum:", max)
}

func convertToBinaryFunc(x int) string {
	var num int = x
	var binary []int

	var result string

	for num != 0 {
		binary = append(binary, num%2)
		num = num / 2
	}

	if len(binary) == 0 {
		fmt.Printf("%d\n", 0)
	} else {
		for i := len(binary) - 1; i >= 0; i-- {
			result += fmt.Sprintf("%d", binary[i])
		}
	}

	return result
}

func convertToBinaryPro(x int) {
	var num int = x
	var binary []int

	for num != 0 {
		binary = append(binary, num%2)
		num = num / 2
	}

	if len(binary) == 0 {
		fmt.Printf("%d\n", 0)
	} else {
		for i := len(binary) - 1; i >= 0; i-- {
			fmt.Print(binary[i])
		}
	}
}
