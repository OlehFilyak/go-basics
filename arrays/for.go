package main

import (
	"fmt"
)

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a) // [1 2 3 4 5]

	// CHANGE LOCAL COPY OF ELEMENT
	for _, elem := range a {
		elem = 100
		fmt.Println(elem)

		// 100
		// 100
		// 100
		// 100
		// 100
	}
	fmt.Println(a) // [1 2 3 4 5]

	// CHANGE ELEMENT OF ARRAY
	for idx := range a {
		a[idx] = 100
		fmt.Println(a[idx])

		// 100
		// 100
		// 100
		// 100
		// 100
	}
	fmt.Println(a) // [100 100 100 100 100]
}
