package main

import (
	"fmt"
)

// Associative arrays are widely used in algorithmic problems.
// When the amount of data is more than a few dozen, searching for a value in a map is more efficient than in an array.
// Based on this information, try solving the following problem, which is often offered at job interviews.
// Given an array of integers A and an integer k. Find and print the indices of a pair of numbers whose sum equals k.
// If there are no such numbers, return an empty slice. The indexes can be returned in any order.

func main() {
	var arr = []int{4, 5, 67, 343, 343, 343, 334, 344, 4342, 324}
	fmt.Println(find(arr, 72))
}

func find(arr []int, k int) []int {
	// make an empty map
	m := make(map[int]int)
	// put array indexes into it, and use the value itself as a key
	for i, a := range arr {
		if j, ok := m[k-a]; ok { // if value k-a is already in the array, it means that arr[j] + arr[i] = k and we have found it
			return []int{i, j}
		}
		// if the value we are looking for is absent, then we add the current index and the value to the map
		m[a] = i
	}
	fmt.Println(m)
	// no pairs of matching numbers found
	return nil
}

// as you can see, the algorithm will pass through the array only once
// if we were to search for a matching value each time we went through the array, we would have to do a lot more calculations
