package main

import (
	"fmt"
)

func main() {
	messages := []string{
		"message 1",
		"message 2",
		"message 3",
		"message 4",
	}

	for _, message := range messages {
		// fmt.Println(index)
		fmt.Println(message)
	}

	// for i := 0; i < len(messages); i++ {
	// 	fmt.Println(i)
	// 	fmt.Println(messages[i])
	// }

}
