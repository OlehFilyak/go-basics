package main

import (
	"fmt"
	"reflect"
)

func main() {
	type Name string
	type Fruit string

	var fruit Fruit
	var name Name

	fruit = "Apple"
	// name = fruit //error
	name = Name(fruit) // now it works

	fmt.Println(fruit)
	fmt.Println(name)
	fmt.Println(reflect.TypeOf(fruit))
	fmt.Println(reflect.TypeOf(name))

	var str string = "Hello, world!"
	println(str[0])
	println(string(str[0]))
}
