package main

import "fmt"

func main() {
	// staticFunc()
	// staticFunc()
	// printMessage("First message")
	// printMessage("Second message")
	message := sayHello("Oleh", 29)
	printMessage(message)
}

// func staticFunc() { // функції дають можливість викликати певну логіку в різних місцях
// 	fmt.Println("function staticFunc")
// }

func printMessage(message string) {
	fmt.Println(message)
}

func sayHello(name string, age int) string {
	return fmt.Sprintf("Привіт, %s! Тобі %d років.", name, age)
}
