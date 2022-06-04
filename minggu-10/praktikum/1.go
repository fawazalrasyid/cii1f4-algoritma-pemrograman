package main

import "fmt"

type set [2022]int

func exist(T set, n int, val int) bool {
	var is_exist bool = false
	for i := 0; i < n; i++ {
		if val == T[i] {
			is_exist = true
		}
	}
	return is_exist
}

func inputSet(T *set, n *int) {
	*n = 1
	fmt.Scan(&T[0])
	fmt.Scan(&T[1])
	for exist(*T, *n, T[*n]) != true {
		*n++
		fmt.Scan(&T[*n])
	}
}

func findIntersection(T1, T2 set, n, m int, T3 *set, h *int) {
	*h = 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if T1[i] == T2[j] {
				T3[*h] = T1[i]
				*h++
			}
		}
	}
}

func printSet(T set, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(T[i], " ")
	}
}

func main() {
	var s1, s2, s3 set
	var n1, n2, n3 int

	inputSet(&s1, &n1)
	inputSet(&s2, &n2)

	findIntersection(s1, s2, n1, n2, &s3, &n3)
	printSet(s3, n3)
}
