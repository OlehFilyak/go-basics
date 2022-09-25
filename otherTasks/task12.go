package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var a int
	fmt.Scan(&a)
	var b int

	var mySlice []int
	for i := 0; i <= 100; i++ {
		b = powInt(2, i)
		if b > a {
			break
		}
		mySlice = append(mySlice, b)
	}

	fmt.Println(strings.Trim(fmt.Sprint(mySlice), "[]"))
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
