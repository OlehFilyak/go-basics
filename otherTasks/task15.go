package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Из натурального числа удалить заданную цифру.
// Входные данные
// Вводятся натуральное число и цифра, которую нужно удалить.
// Выходные данные
// Вывести число без заданных цифр.
// Sample Input:
// 38012732
// 3
// Sample Output:
// 801272

func main() {
	var a int
	fmt.Scan(&a)
	var b int
	fmt.Scan(&b)
	var c = strconv.Itoa(a)
	var d = strings.ReplaceAll(c, strconv.Itoa(b), "")
	fmt.Println(d)
}
