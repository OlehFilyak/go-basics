package main

import (
	"fmt"
)

// Количество нулей
// По данным числам, определите количество чисел, которые равны нулю.
// Входные данные
// Вводится натуральное число N, а затем N чисел.
// Выходные данные
// Выведите количество чисел, которые равны нулю.
// Sample Input:
// 5
// 1
// 8
// 100
// 0
// 12
// Sample Output:
// 1

func main() {
	var n int
	fmt.Scan(&n)
	var b int
	var counter int

	for i := 0; i < n; i++ {
		fmt.Scan(&b)
		if b == 0 {
			counter++
		}
	}
	fmt.Println(counter)
}
