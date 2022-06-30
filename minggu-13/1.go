package main

// func searchFirstIndexOdd(A []int, n int) int {
// 	if n == 0 {
// 		return -1
// 	} else if A[n-1]%2 == 1 {
// 		return n - 1
// 	} else {
// 		return searchFirstIndexOdd(A, n-1)
// 	}
// }

func firstIndexOddArrayRecursive(A []int, n int) int {
	if n == 0 {
		return -1
	} else if A[n-1]%2 == 1 {
		return n - 1
	} else {
		return firstIndexOddArrayRecursive(A, n-1)
	}
}

func cariIndexArrayGanjil(A []int, n int) int {
	if n == 0 {
		return -1
	} else if A[n-1]%2 == 1 {
		return n - 1
	} else {
		return cariIndexArrayGanjil(A, n-1)
	}
}