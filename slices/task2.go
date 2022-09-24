package main

import (
	"fmt"
)

// На ввод подаются пять целых чисел, которые записываются в массив. Однако эта часть программы уже написана.
// Вам нужно написать фрагмент кода, с помощью которого можно найти и вывести максимальное число в этом массиве.
// Sample Input:
// 2
// 3
// 56
// 45
// 21
// Sample Output:
// 56

func main() {

	array := []int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}

	var b int
	b = array[0]
	// fmt.Println(a)
	for i := 0; i < len(array); i++ {
		if b < array[i] {
			b = array[i]
		}
	}
	fmt.Println(b)
}
