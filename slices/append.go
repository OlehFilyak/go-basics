package main

import (
	"fmt"
)

func main() {

	// Встроенные функции для работы со срезами
	// Для работы со срезами в Go предусмотрены встроенные функции append и copy.

	// Функция append - добавление элементов в срез
	// Функция append позволяет добавить в срез новый элемент, выглядит она следующим образом:

	// func append(slice []Type, elems ...Type) []Type
	// В качестве первого аргумента функция получает срез, в который необходимо добавить новые элементы,
	// второй и последующий элементы - это элементы совместимого со срезом типа, которые необходимо добавить в срез.
	// Функция возвращает новый срез, содержащий ранее содержавшиеся в срезе элементы,
	// а также новые элементы, переданные в качестве аргумента функции append.
	// Подробно функции будут рассмотрены в разделе 2.1. Функции.

	a := []int{1, 2, 3}
	a = append(a, 4, 5)

	fmt.Println(a) // [1 2 3 4 5]
}
