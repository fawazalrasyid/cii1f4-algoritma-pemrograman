package main

import "fmt"

func main() {
	fmt.Println(f(8))
}

func f(n int) int {
	if n >= 1 && n <= 4 {
		return 1
	} else {
		return f(n-1) + f(n-2) + f(n-3) + f(n-4)
	}

}
