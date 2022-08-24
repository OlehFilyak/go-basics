package main

import "fmt"

var a, b, c int // доступні в області бачення поточного пакету

func main() {
	a, b, c := 1, 2, 3 // local variables
	fmt.Println(a, b, c)

	print()
	any()
}

func print() {
	a, b, c := 4, 5, 6
	fmt.Println(a, b, c)
}

func any() {
	fmt.Println(a, b, c)
}
