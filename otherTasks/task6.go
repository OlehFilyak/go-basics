package main

import (
	"fmt"
)

// Даны два числа. Найти их среднее арифметическое.
// Формат входных данных
// На вход дается два целых положительных числа a и b.
// Формат выходных данных
// Программа должна вывести среднее арифметическое чисел a и b (ответ может быть целым числом или дробным)
// Sample Input 1:
// 3 5
// Sample Output 1:
// 4
// Sample Input 2:
// 277 154
// Sample Output 2:
// 215.5

func main() {
	var a float64
	var b float64

	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Println((a + b) / 2)
}
