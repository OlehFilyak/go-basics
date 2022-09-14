package main

import (
	"fmt"
)

func main() {

	// DECLARATION
	var arrayWithZeroValues [3]int
	fmt.Println("arrayWithZeroValues is", arrayWithZeroValues)

	// ARRAY LITERAL
	var arrayLiteral = [3]int{10, 20, 30}
	fmt.Println("arrayLiteral is", arrayLiteral)
	var arrayLiteral1 = [...]int{10, 20, 30}
	fmt.Println("arrayLiteral == arrayLiteral1?", arrayLiteral == arrayLiteral1)

	// CHANGE THE VALUE AND READ THE VALUE
	arrayLiteral[1] = 99
	fmt.Println(arrayLiteral[1])
	fmt.Println(arrayLiteral[2])

	// SPARSE ARRAY
	var sparseArray = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println("sparseArray is", sparseArray)

	//MULTIDEMENSIONAL ARRAYS
	var multidemensionalArrays [2][3]int
	fmt.Println("multidemensionalArrays is", multidemensionalArrays)

}
