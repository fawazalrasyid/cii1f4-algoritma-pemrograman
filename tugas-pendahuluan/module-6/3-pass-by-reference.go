package main

import "fmt"

func main() {

	var a string = "nasi goreng"
	var b string = "bakso"

	tukar(&a, &b)
	fmt.Println(a, b)

}

func tukar(a, b *string) {
	var temp string

	temp = *a
	*a = *b
	*b = temp
}
