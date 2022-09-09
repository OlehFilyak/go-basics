package main

import "fmt"

var a, b, c int // variables are available in the viewport of the current package

func main() {
	i := 10
	if i == 10 {
		// change the value of the variable i
		i += 5
		if i == 15 {
			// this block creates a new variable i, which
			// overrides the variable with the same name defined above
			// such situations should be avoided in practice
			i := 7
			fmt.Println(i)
			// the scope of this variable is limited to the block
		}
	}
	fmt.Println(i)

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
