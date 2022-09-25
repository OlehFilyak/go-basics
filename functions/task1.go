package main

import "fmt"

// Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.
// Входные данные
// Вводится четыре числа.
// Выходные данные
// Необходимо вернуть из функции наименьшее из 4-х данных чисел
// Sample Input:
// 4 5 6 7
// Sample Output:
// 4

func main() {
	minimumFromFour()
}

func minimumFromFour() int {
	//print your code here
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	var min int
	if a <= b && a <= c && a <= d {
		return a
	}
	if b <= a && b <= c && b <= d {
		return b
	}
	if c <= a && c <= b && c <= d {
		return c
	}
	if d <= a && d <= c && d <= b {
		return d
	}

	return min
}
