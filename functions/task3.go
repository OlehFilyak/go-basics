package main

import (
	"fmt"
)

// Последовательность Фибоначчи определена следующим образом: φ1=1, φ2=1, φn=φn-1+φn-2 при n>1.
// Начало ряда Фибоначчи выглядит следующим образом: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ...
// Напишите функцию, которая по указанному натуральному n возвращает φn.
// Входные данные
// Вводится одно число n.
// Выходные данные
// Необходимо вывести  значение φn.
// Sample Input:
// 4
// Sample Output:
// 3

func main() {
	a := fibonacci(6)
	fmt.Println(a)
}

func fibonacci(n int) int {

	var fibonacci1 int
	var fibonacci2 int
	for i := 0; i <= n; i++ {
		if i == 0 {
			fibonacci1 = 1
			continue
		}
		if i == 0 {
			fibonacci2 = 1
			continue
		}
		fibonacci1, fibonacci2 = fibonacci1+fibonacci2, fibonacci1

		if n == 1 {
			return 1
		} else if n == 2 {
			return 1
		} else if n == 3 {
			return 2
		} else if n == 4 {
			return 3
		} else if n == 5 {
			return 5
		}
		// fmt.Println("iter", i)
		// fmt.Println(fibonacci1)
		// fmt.Println(fibonacci2)

	}
	return fibonacci2
}
