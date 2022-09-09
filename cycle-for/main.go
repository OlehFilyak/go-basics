package main

import (
	"fmt"
)

func main() {
	// INFINITE LOOP

	// for {
	// 	// code executed inside an infinite loop
	// }

	// for ;; {}

	// for ; true; {}

	// TRHEE-PART CYCLE
	// create a variable
	v := 0
	//
	for i := 1; i < 10; i++ {
		// increment the variable
		v++
	}
	// display the result of the screen
	fmt.Println(v)

	// COMPLEX VIEW OF THE CYCLE
	// for a, b := 5, 10; a < 10 && b < 20; a, b = a+1, b+2 {
	// 	// do stuff
	// }

	// WHILE CYCLE
	// // create a variable
	// i := 0
	// // describe the precondition
	// for i < 5 {
	// 	// increment the variable
	// 	i++
	// }
	// // display the result of the screen
	// fmt.Println(i)

	// RANGE CYCLE
	// // create an array
	// array := [3]int{1, 2, 3}
	// // iterating
	// for arrayIndex, arrayValue := range array {
	// 	fmt.Printf("array[%d]: %d\n", arrayIndex, arrayValue)
	// }

	// BREAK AND CONTINUE KEY WORDS
	// sum, limit := 0, 100
	// for i := 0; true; i++ {
	// 	if i%2 != 0 {
	// 		continue // moving to the next number, since i is odd
	// 	}

	// 	if sum+i > limit {
	// 		break // exit the loop, because the sum will exceed the given limit
	// 	}

	// 	sum += i
	// }
	// fmt.Println(sum)
}
