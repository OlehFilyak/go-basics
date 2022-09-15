package main

import (
	"fmt"
)

// Требуется написать программу, при выполнении которой с клавиатуры
// считываются два натуральных числа A и B (каждое не более 100, A < B).
// Вывести сумму всех чисел от A до B  включительно.

// Sample Input:
// 1 5
// Sample Output:
// 15

func main() {
	var a int
	fmt.Scan(&a)

	var b int
	fmt.Scan(&b)

	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Println(sum)
}
