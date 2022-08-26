package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMin(13333, 1, 23, 4, -665, 666, 432))
}

func findMin(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}

	min := numbers[0]

	for _, i := range numbers {
		if i < min {
			min = i
		}
	}

	return min
}
