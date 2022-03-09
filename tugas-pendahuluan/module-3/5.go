package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	consoleReader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name : ")

	var nama, _ = consoleReader.ReadString('\n')

	fmt.Println("Your name is : ", nama)

}
