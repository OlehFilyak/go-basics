package main

import (
	"fmt"
)

func main() {

	var a []int
	var b []int = []int{1, 2, 3}
	c := []int{1, 2, 3}
	d := []int{1: 12}

	fmt.Println(a) // []
	fmt.Println(b) // [1 2 3]
	fmt.Println(c) // [1 2 3]
	fmt.Println(d) // [0 12]

	x := make([]int, 10, 10) // [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(x)
}
