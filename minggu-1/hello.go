package main    // header

import "fmt"    // library fmt

func main() {   // function
    // cara 1 untuk membuat variabel
    var s string
    var a, b int

    // input
    fmt.Scan(&s, &a, &b)

    // cara 2 untuk membuat variabel / shorthand
    hasil_penjumlahan := a+b

    // print
    fmt.Println("Kata =", s)
    fmt.Println("Jumlah =", hasil_penjumlahan)

}