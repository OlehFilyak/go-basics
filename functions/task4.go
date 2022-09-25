package main

import (
	"fmt"
)

// Напишите функцию sumInt, принимающую переменное количество аргументов типа int,
// и возвращающую количество полученных функцией аргументов и их сумму.
// Пакет "fmt" уже импортирован, функция и пакет main объявлены.
// Пример вызова вашей функции:
// a, b := sumInt(1, 0)
// fmt.Println(a, b)
// Результат: 2, 1
// Sample Input:
// Sample Output:

func main() {
	a, b := sumInt(1, 0)
	fmt.Println(a, b)
}

/* ЗАДАНИЕ:
 * Напишите функцию sumInt, получающую переменное число аргументов типа int,
 * и возвращающую количество переданных аргументов и их сумму.
 */
func sumInt(a ...int) (int, int) {
	var counter int
	var sum int
	for _, elem := range a {
		sum += elem
		counter++
	}
	return counter, sum
}
