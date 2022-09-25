package main

import (
	"fmt"
)

// Входные данные
// Даны три натуральных числа a, b, c. Определите, существует ли треугольник с такими сторонами.
// Выходные данные
// Если треугольник существует, выведите строку "Существует", иначе выведите строку "Не существует". Строку выводите без кавычек.
// Sample Input:
// 4 5 6
// Sample Output:
// Существует

func main() {
	var a int
	var b int
	var c int

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	switch {
	case (a+b > c && c >= a && c >= b) || (a+c > b && b >= a && b >= c) ||
		(b+c > a && a >= c && a >= b):
		fmt.Println("Существует")
	default:
		fmt.Println("Не существует")
	}
}
