package main

import (
	"fmt"
)

// Количество минимумов
// Найдите количество минимальных элементов в последовательности.
// Входные данные
// Вводится натуральное число N, а затем N целых чисел последовательности.
// Выходные данные
// Выведите количество минимальных элементов последовательности.
// Sample Input:
// 3
// 21
// 11
// 4
// Sample Output:
// 1

func main() {
	var n int
	fmt.Scan(&n)
	var b int
	var min int
	var counter int
	for i := 0; i < n; i++ {

		fmt.Scan(&b)
		if i == 0 {
			min = b
			counter = 1
			continue
		}
		if b == min {
			counter++
		} else if b < min {
			counter = 1
			min = b
		} else if b > min {
			continue
		}
	}

	fmt.Println(counter)
}
