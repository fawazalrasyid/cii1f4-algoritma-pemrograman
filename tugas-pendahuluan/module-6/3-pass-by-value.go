package main

import "fmt"

func main() {

	var a string = "nasi goreng"
	var b string = "bakso"

	tuker(a, b)
	fmt.Println(a, b)

}

func tuker(a, b string) {
	var temp string

	temp = a
	a = b
	b = temp
}
