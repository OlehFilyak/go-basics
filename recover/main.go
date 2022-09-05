package main

import (
	"fmt"
)

func main() {
	defer handlerPanic()

	messages := []string{
		"message 1",
		"message 2",
		"message 3",
		"message 4",
	}
	messages[4] = "messages 5"

	fmt.Println(messages) // не виконається
}

func handlerPanic() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
	fmt.Println("handlerPanic() виконалася успішно")
}
