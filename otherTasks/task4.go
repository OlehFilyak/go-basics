package main

import (
	"fmt"
)

// Заданы три числа - a, b, c (a < b < c)a,b,c(a<b<c) - длины сторон треугольника.
// Нужно проверить, является ли треугольник прямоугольным.
// Если является, вывести "Прямоугольный". Иначе вывести "Непрямоугольный"
// Sample Input:
// 6 8 10
// Sample Output:
// Прямоугольный

func main() {
	var a int
	var b int
	var c int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	var gip int = c
	if gip*gip == b*b+a*a {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}

}
