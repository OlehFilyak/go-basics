package main

import (
	"fmt"
)

// Вклад в банке составляет x рублей. Ежегодно он увеличивается на p процентов, после чего дробная часть копеек отбрасывается.
// Каждый год сумма вклада становится больше.
// Определите, через сколько лет вклад составит не менее y рублей.
// Входные данные
// Программа получает на вход три натуральных числа: x, p, y.
// Выходные данные
// Программа должна вывести одно целое число.
// Sample Input:
// 100
// 10
// 200
// Sample Output:
// 8

func main() {
	var x int
	var p int
	var y int

	fmt.Scan(&x)
	fmt.Scan(&p)
	fmt.Scan(&y)

	x = x * 100
	y = y * 100

	i := 0
	for x <= y {
		i++
		x = x + (x * p / 100)
	}
	fmt.Println(i)
}