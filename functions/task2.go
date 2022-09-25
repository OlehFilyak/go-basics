package main

// import "fmt"

// Напишите "функцию голосования", возвращающую то значение (0 или 1), которое среди значений ее аргументов x, y, z встречается чаще.
// Входные данные
// Вводится 3 числа - x, y и z (x, y и z равны 0 или 1).
// Выходные данные
// Необходимо вернуть значение функции от x, y и z.
// Sample Input:
// 0 0 1
// Sample Output:
// 0
func main() {
	vote(0, 1, 1)
}

func vote(x int, y int, z int) int {
	//print your code here
	var counter0 int
	var counter1 int
	if x == 0 {
		counter0++
	} else {
		counter1++
	}
	if y == 0 {
		counter0++
	} else {
		counter1++
	}
	if z == 0 {
		counter0++
	} else {
		counter1++
	}
	if counter1 > counter0 {
		return 1
	} else {
		return 0
	}
	return 0
}
