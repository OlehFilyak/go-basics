package main

import (
	"fmt"
)

// var count int

func main() {
	inc := increment()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())

	// fmt.Println(increment2())
	// fmt.Println(count)
	// fmt.Println(increment2())
	// fmt.Println(count)
	// fmt.Println(increment2())
	// fmt.Println(count)
}

func increment() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// func increment2() int {
// 	// count := 0
// 	count++
// 	fmt.Println("Inside increment2", count)
// 	return count
// }
