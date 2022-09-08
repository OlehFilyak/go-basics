package main

import (
	"fmt"
	"reflect"
)

func main() {
	// type rune uint32
	// type byte uint8

	type Name string
	type Fruit string

	var fruit Fruit
	var name Name

	fruit = "Apple"
	fmt.Println(reflect.TypeOf(fruit))
	fmt.Println(reflect.TypeOf(name))
	// name = fruit // cannot use fruit (variable of type Fruit) as type Name in assignment
	//
}
