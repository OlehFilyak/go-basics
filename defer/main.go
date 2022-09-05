package main

import (
	"fmt"
)

func main() {
	defer printMessage()

	fmt.Println("main() 1")
	fmt.Println("main() 2")
	fmt.Println("main() 3")
}

func printMessage() {
	fmt.Println("printMessage()")
}
