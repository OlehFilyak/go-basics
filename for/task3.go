package main

import (
	"fmt"
)

// Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8.
// Программа в первой строке получает на вход число n - количество чисел в последовательности,
// во второй строке -- n чисел, входящих в данную последовательность.

// Sample Input:
// 5
// 38 24 800 8 16
// Sample Output:
// 40

func main() {
	var a int
	fmt.Scan(&a)

	var b int = 0
	sum := 0
	for i := 1; i <= a; i++ {
		// fmt.Println(i)
		fmt.Scan(&b)

		if b <= 99 && b >= 10 && b%8 == 0 {
			sum += b
		}
	}
	fmt.Println(sum)
}
