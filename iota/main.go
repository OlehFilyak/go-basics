package main

import (
	"fmt"
)

const (
	Black = iota
	Gray
	White
)

// the counter is reset
const (
	Yellow = iota
	Red
	Green = iota // this assignment will not reset iota
	Blue
)

const (
	_ = iota * 10 // note that it is possible to skip constants
	ten
	hundred
	thousand
)

const (
	hello = "Hello, world!" // iota is 0
	one1  = 1               // iota is 1
	black = iota            // iota is 2
	gray
)

const (
	one = iota + iota + 1 // need 1, 3, 5, 7, 9, 11
	three
	five
	seven
	nine
	eleven
)

// WEEKDAY
type Weekday int

const (
	Monday Weekday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func NextDay(day Weekday) Weekday {
	return (day % 7) + 1
}

func main() {
	fmt.Println(Black, Gray, White)
	fmt.Println(Yellow, Red, Green, Blue)

	fmt.Println(ten, hundred, thousand)
	fmt.Println(hello, one1, black, gray)

	fmt.Println(one, three, five, seven, nine, eleven)

	var today Weekday = Sunday
	tomorrow := NextDay(today)
	fmt.Println("today =", today, "tomorrow =", tomorrow)
}
