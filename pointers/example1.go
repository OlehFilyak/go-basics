package main

import (
	"fmt"
)

func main() {
	a := 200
	b := &a
	*b++
	// a := 200
	// var b *int = &a
	fmt.Println(a)
}
