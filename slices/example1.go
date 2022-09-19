package main

import (
	"errors"
	"fmt"
)

func main() {

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

	printMessagesSlices(messagesSlice)
	fmt.Println(messagesSlice)

	// weekTempArr := [7]int{1, 2, 3, 4, 5, 6, 7}
	// workDaysSlice := weekTempArr[:5]
	// weekendSlice := weekTempArr[5:]
	// fromTuesdayToThursDaySlice := weekTempArr[1:4]
	// weekTempSlice := weekTempArr[:]

	// fmt.Println(workDaysSlice, len(workDaysSlice), cap(workDaysSlice))                                        // [1 2 3 4 5] 5 7
	// fmt.Println(weekendSlice, len(weekendSlice), cap(weekendSlice))                                           // [6 7] 2 2
	// fmt.Println(fromTuesdayToThursDaySlice, len(fromTuesdayToThursDaySlice), cap(fromTuesdayToThursDaySlice)) // [2 3 4] 3 6
	// fmt.Println(weekTempSlice, len(weekTempSlice), cap(weekTempSlice))                                        // [1 2 3 4 5 6 7] 7 7
}

func printMessagesSlices(messagesSlice []string) error {
	if len(messagesSlice) == 0 {
		return errors.New("empty array")
	}
	messagesSlice[1] = "111111111"
	fmt.Println(messagesSlice)
	return nil
}
