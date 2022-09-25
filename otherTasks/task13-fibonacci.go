package main

import (
	"fmt"
)

// Номер числа Фибоначчи
// Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является, то есть выведите такое число n, что φn=A.
// Если А не является числом Фибоначчи, выведите число -1.
// Входные данные
// Вводится натуральное число.
// Выходные данные
// Выведите ответ на задачу.
// Sample Input:
// 8
// Sample Output:
// 6

func main() {
	var a int
	fmt.Scan(&a)

	var fibonacci1 int
	var fibonacci2 int
	for i := 0; i <= a; i++ {
		if i == 0 {
			fibonacci1 = 0
			continue
		}
		if i == 1 {
			fibonacci2 = 1
			continue
		}
		fibonacci1, fibonacci2 = fibonacci1+fibonacci2, fibonacci1
		if fibonacci1 == a {
			fmt.Println(i - 1)
			return
		} else if a == 2 {
			fmt.Println(3)
			return
		} else if a == 3 {
			fmt.Println(4)
			return
		} else if a == 5 {
			fmt.Println(5)
			return
		}

	}

	fmt.Println(-1)

}
