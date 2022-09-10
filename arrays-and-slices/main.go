package main

import (
	"errors"
	"fmt"
)

func main() {
	// thisWeekTemp := [7]int {-3,5,7} // [-3 5 7 0 0 0 0]
	// rgbColor := [3]uint8 {255,255,128} // [255 255 128]

	// rgbColor := [...]uint8{255,255,128} // [255 255 128] len = 3
	// rgbaColor := [...]uint8{255,255,128,1} // [255 255 128 1] len = 4

	// thisWeekTemp := [7]int{6: 11, 2: 3} // [0 0 3 0 0 0 0 11]
	// fmt.Println(thisWeekTemp)

	// // Multivariate arrays
	// var thisMonthTemp [4][7]int
	// fmt.Println(thisMonthTemp)

	//TASK calculate the average temperatures for the week
	// var weekTemp = [7]int{5, 4, 6, 8, 11, 9, 5}

	// sumTemp := 0

	// for i := 0; i < len(weekTemp); i++ {
	// 	sumTemp += weekTemp[i]
	// }

	// average := sumTemp / len(weekTemp)

	var weekTemp = [7]int{5, 4, 6, 8, 11, 9, 5}

	sumTemp := 0

	for _, temp := range weekTemp {
		sumTemp += temp
	}

	average := sumTemp / len(weekTemp)
	fmt.Println(average)

	// change all elements of array to zero
	// var weekTemp = [7]int{5, 4, 6, 8, 11, 9, 5}
	// // weekTemp [5 4 6 8 11 9 5]
	// for _, temp := range weekTemp {
	// temp = 0
	// }
	// // weekTemp [5 4 6 8 11 9 5] ! - values have not changed
	// // if the value of the element is not used, you can omit the second variable
	// for i := range weekTemp {
	// weekTemp[i] = 0
	// }
	// // weekTemp [0 0 0 0 0 0] ! - values have changed

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
