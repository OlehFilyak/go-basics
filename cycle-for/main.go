package main

import (
	"fmt"
)

func main() {
	// for x := 0; x < 10; x++ {
	// 	fmt.Println(x)
	// }

	counter := 0
	for {
		if counter == 100 {
			break
		}
		counter++
		fmt.Println(counter)
	}
}
