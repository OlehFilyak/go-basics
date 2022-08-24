package main

import (
	"fmt"
	"reflect"
)

func main() {
	message := "Hello world"
	var firstName string = "Oleh"
	const birthadayYear int = 1992

	fmt.Println(message)
	fmt.Println(firstName)
	fmt.Println(birthadayYear)

	fmt.Println(reflect.TypeOf(message))
	fmt.Println(reflect.TypeOf(firstName))
	fmt.Println(reflect.TypeOf(birthadayYear))
}
