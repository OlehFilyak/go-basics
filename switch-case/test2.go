package main

import "fmt"

// На ввод подается целое число.
// Если число положительное - вывести сообщение "Число положительное", если число отрицательное - "Число отрицательное".
// Если подается ноль - вывести сообщение "Ноль". Выводить сообщение без кавычек.

func main() {
	var a int
	fmt.Scan(&a)

	switch {
	case a == 0:
		fmt.Println("Ноль")
	case a > 0:
		fmt.Println("Число положительное")
	case a < 0:
		fmt.Println("Число отрицательное")
	}
}