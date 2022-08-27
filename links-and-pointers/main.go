package main

import (
	"fmt"
)

func main() {
	message := "I'll become golang developer soon."

	fmt.Println(message)
	ChangeMessage(&message) // we pass a reference to an area in memory
	fmt.Println(message)
	fmt.Println(&message) // the area in memory is displayed

	dontChangeMessage(message)
	fmt.Println(message)

	number := 7
	var p *int //an empty pointer
	fmt.Println(p)
	p = &number
	fmt.Println(p)
	fmt.Println(&number)
	*p = 10
	fmt.Println(number)
}

func ChangeMessage(message *string) {
	*message += " Some addition from func addToMessage" // we refer to the value by the area in memory
}

func dontChangeMessage(message string) {
	message += " Some addition from func dontChangeMessage"
	fmt.Println(&message)
}
