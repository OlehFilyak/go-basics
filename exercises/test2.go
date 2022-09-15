package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a) // считаем переменную 'a' с консоли
	b := a % 30
	c := (a - b) / 30
	d := (a % 30) * 2

	// fmt.Println(b)
	fmt.Println("It is", c, "hours", d, "minutes.")
}
