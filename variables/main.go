package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	// SHORT DECLARATION - ONLY WHEN YOU WITHIN IN FUNC
	shortDeclaration := "shortDeclaration"
	fmt.Println("shortDeclaration is", shortDeclaration)
	// ASSIGN VALUE TO EXICT VARIABLE
	// x := 10
	// x, y := 30, "hello"

	// EXPLICIT DECLARATION
	var explicitDeclaration string = "explicitDeclaration"
	fmt.Println("explicitDeclaration is", explicitDeclaration)

	// DECLARATION WITH DEFAULT TYPE
	var declarationWithDefaultType = "declarationWithDefaultType"
	fmt.Println("declarationWithDefaultType is", declarationWithDefaultType)

	// DECLARE AND ASSIGN IT THE ZERO VALUE
	var declarationWithZeroValue string
	fmt.Println("declarationWithZeroValue is", declarationWithZeroValue)

	// DECLARE MULTIPLE VARIABLES
	// SAME TYPE
	var a, b int
	a, b = 2, 1 // a == 5, b == 10
	a, b = b, a // swap: a == 10, b == 5
	fmt.Println("a is", a)
	fmt.Println("b is", b)

	var c, d int // Multiple assignment
	c, d = 3, 4  // The value is assigned in order: a == 5 и b == 10
	fmt.Println("declarationWithZeroValue is", declarationWithZeroValue)
	fmt.Println("c is", c)
	fmt.Println("d is", d)

	// DIFFERENT TYPE
	var e, f = 5, "banana"
	fmt.Println("declarationWithZeroValue is", declarationWithZeroValue)
	fmt.Println("e is", e)
	fmt.Println("f is", f)

	// DECLARATION LIST
	var (
		g    int
		h        = 20
		j    int = 30
		k, o     = 40, "hello"
		p, q string
	)
	fmt.Println("g is", g)
	fmt.Println("h is", h)
	fmt.Println("j is", j)
	fmt.Println("k is", k)
	fmt.Println("o is", o)
	fmt.Println("p is", p)
	fmt.Println("q is", q)

	// CHANGE DEFAULT TYPE
	int64Var := int64(5) // var int64Var int64 = 5
	// var int64Var int64 = 5 // idiomatic
	float32Var := float32(101.3) // var floatVar float32 = 101.3

	pi, ex := 3.1415, 2.7183
	pi, ex = 3.14159, 2.71828

	df, err := os.Open("myfile.txt")

	fmt.Println(reflect.TypeOf(int64Var))
	fmt.Println(reflect.TypeOf(float32Var))
	fmt.Println(reflect.TypeOf(pi))
	fmt.Println(reflect.TypeOf(ex))
	fmt.Println(reflect.TypeOf(df))
	fmt.Println(reflect.TypeOf(err))

	fmt.Println(a, b, c, d)

	// CONST
	const birthadayYear int = 1992
	fmt.Println(birthadayYear)
	fmt.Println(reflect.TypeOf(birthadayYear))

	const id = 100      // untyped int
	const typedId = 100 // «This constant can only be assigned directly to an int.»
	fmt.Println(reflect.TypeOf(id))
	fmt.Println(reflect.TypeOf(typedId))

	// ASSIGN UNTYPED CONST
	var y int64 = id
	var z float64 = id

	fmt.Println("y=", y, "z=", z)
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
