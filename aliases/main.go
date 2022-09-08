package main

import (
	"fmt"
)

func main() {
	type MyString = string // MyString here is an alias of the string type

	var a string // you can use one of these conditions to check for an empty string
	var b MyString
	a = b // there is no error
	fmt.Println(a == b)
}
