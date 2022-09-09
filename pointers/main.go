package main

import (
	"fmt"
)

func main() {
	var b *int    // Syntax
	var a int = 5 // pointer to var
	p := &a
	fmt.Println(a, p, b)

	const c = 5
	// p1 := &"abc" // compilation error
	// p2 := &с // compilation error
	fmt.Println(c)

	var p1 *int
	var t int = 5
	var r string = "abc"
	p1 = &t
	// p1 = &r // compilation error
	fmt.Println(p1)
	fmt.Println(r)

	type A struct {
		IntField int
	}
	// The A{} literal creates a variable of type A in memory. Then it takes a pointer from it
	g := &A{
		IntField: 10,
	}
	fmt.Println(g)

	// g := new(A) //  the same as &A{}

	i := 42
	o := &i
	fmt.Println(*p) // read the value of variable i via pointer o
	*o = 21         // write value 21 to the i variable via pointer o
	fmt.Println(o)

	// pointer to pointer
	fmt.Println(*o)

	// pointer to pointer <nil>
	// var q *int
	// fmt.Println(*q) // panic error

}
