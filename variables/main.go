package main

import (
	"fmt"
	"reflect"
)

func main() {
	message := "Hello world"
	var firstName string = "Oleh"
	const birthadayYear int = 1992

	fmt.Println(message)
	fmt.Println(firstName)
	fmt.Println(birthadayYear)

	fmt.Println(reflect.TypeOf(message))
	fmt.Println(reflect.TypeOf(firstName))
	fmt.Println(reflect.TypeOf(birthadayYear))

	var c, d int // Multiple assignment
	c, d = 3, 4  // The value is assigned in order: a == 5 и b == 10

	var a, b int
	a, b = 2, 1 // a == 5, b == 10
	a, b = b, a // swap: a == 10, b == 5

	fmt.Println(a, b, c, d)
}
