package main

import (
	"fmt"
)

func main() {

	result := Cube(5) // вызов функции
	fmt.Println(result)
}

func Cube(x int) int { // декларация функции
	return x * x * x // тело функции
}

// func MyFunction(arg1 arg1type, arg2 arg2type) resultType{
//     // тело функции
// }

func Index(st string, a rune) (index int, ok bool) {
	for i, c := range st {
		if c == a {
			return i, true
		}
	}
	return // вернутся значения по умолчанию
}
