package main

import (
	"fmt"
)

func main() {
	var a float64 = 100.123456
	fmt.Printf("это число %f типа %T \n", a, a)
	// вывод: это число 100.123456 типа float64

	var a1 byte = 's'
	var a2 int = 1234
	fmt.Printf("%q %b \n", a1, a2)
	// вывод: 's' 10011010010

	// использование \n позволяет сделать перенос строки
	var a3 string = "123"
	var a4 string = "1234"
	fmt.Printf("%q \n%s", a3, a4)
	// вывод:
	// "123"
	// 1234
}
