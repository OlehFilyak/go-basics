package main

import "fmt"

// Дано неотрицательное целое число. Найдите и выведите первую цифру числа.

// Формат входных данных
// На вход дается натуральное число, не превосходящее 10000.

// Формат выходных данных
// Выведите одно целое число - первую цифру заданного числа.

// Sample Input:
// 1234
// Sample Output:
// 1

func main() {
	var a int
	fmt.Scan(&a)

	b := a % 10
	// fmt.Println(b)
	c := (a - b) / 10 % 10
	// fmt.Println(c)
	d := (a - b - (c * 10)) / 100 % 10
	// fmt.Println(d)
	e := (a - b - (c * 10) - (d * 100)) / 1000 % 10
	// fmt.Println(e)
	f := (a - b - (c * 10) - (d * 100) - (e * 1000)) / 10000 % 10
	// fmt.Println(f)
	switch {
	case a < 10:
		fmt.Println(b)
	case a < 100:
		fmt.Println(c)
	case a < 1000:
		fmt.Println(d)
	case a < 10000:
		fmt.Println(e)
	case a <= 100000:
		fmt.Println(f)
	default:
		fmt.Println("Uncorrect start value")
	}
}
