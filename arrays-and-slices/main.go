package main

import (
	"errors"
	"fmt"
)

func main() {
	messagesArray := [3]string{"1", "2", "3"}
	empty := [3]int{}
	messagesSlice := []string{"4", "5", "6"}
	var emptySlice []string
	emptySlice5Len := make([]string, 5)

	fmt.Println(emptySlice5Len)
	emptySlice5Len[4] = "44444"
	fmt.Println(emptySlice5Len)
	emptySlice5Len = append(emptySlice5Len, "321")
	fmt.Println(emptySlice5Len)
	fmt.Println(len(emptySlice5Len))
	fmt.Println(cap(emptySlice5Len))
	// emptySlice[0] = "2" // error
	fmt.Println(emptySlice)
	fmt.Println(len(emptySlice))
	// fmt.Println(cap(emptySlice))

	fmt.Println(messagesSlice)

	fmt.Println(messagesArray)
	fmt.Println(messagesArray[1])
	messagesArray[1] = "777"
	fmt.Println(messagesArray)
	fmt.Println(empty)
	printMessagesArray(messagesArray)
	fmt.Println(messagesArray)

	printMessagesSlices(messagesSlice)
	fmt.Println(messagesSlice)
}

func printMessagesArray(messagesArrayCopy [3]string) error {
	if len(messagesArrayCopy) == 0 {
		return errors.New("empty array")
	}
	messagesArrayCopy[1] = "999"
	fmt.Println(messagesArrayCopy)
	return nil
}

func printMessagesSlices(messagesSlice []string) error {
	if len(messagesSlice) == 0 {
		return errors.New("empty array")
	}
	messagesSlice[1] = "111111111"
	fmt.Println(messagesSlice)
	return nil
}
