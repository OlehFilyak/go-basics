package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	message := "Hello world"
	var firstName string = "Oleh"

	fmt.Println(message)
	fmt.Println(firstName)

	int64Var := int64(5)         // var int64Var int64 = 5
	float32Var := float32(101.3) // var floatVar float32 = 101.3

	pi, e := 3.1415, 2.7183
	pi, e = 3.14159, 2.71828

	f, err := os.Open("myfile.txt")

	fmt.Println(reflect.TypeOf(message))
	fmt.Println(reflect.TypeOf(firstName))
	fmt.Println(reflect.TypeOf(int64Var))
	fmt.Println(reflect.TypeOf(float32Var))
	fmt.Println(reflect.TypeOf(pi))
	fmt.Println(reflect.TypeOf(e))
	fmt.Println(reflect.TypeOf(f))
	fmt.Println(reflect.TypeOf(err))

	var c, d int // Multiple assignment
	c, d = 3, 4  // The value is assigned in order: a == 5 и b == 10

	var a, b int
	a, b = 2, 1 // a == 5, b == 10
	a, b = b, a // swap: a == 10, b == 5

	fmt.Println(a, b, c, d)

	// CONST
	const birthadayYear int = 1992
	fmt.Println(birthadayYear)
	fmt.Println(reflect.TypeOf(birthadayYear))

	const id = 100 // untyped int
	fmt.Println(reflect.TypeOf(id))
	var y int64 = id
	var z float64 = id

	fmt.Println("i=", y, "f=", z)
	fmt.Println(reflect.TypeOf(y))
	fmt.Println(reflect.TypeOf(z))

	var qwerty float64
	const intConst = 5
	const floatConst = 5.0
	qwerty = intConst + floatConst // untyped int, untyped float
	fmt.Println(qwerty)

	const (
		pi1      = 3.1415
		e1       // 3.1415
		name     = "John Doe"
		fullName // "John Doe"
	)

	fmt.Println("pi1 =", pi1, "e1 =", e1)
	fmt.Println("name =", name, "fullName =", fullName)
}
