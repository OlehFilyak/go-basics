package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	message, err := dailyWish("thursday")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	// declaration of a local variable in the middle of the block
	i := 6
	switch j := i % 5; {
	case j == 0:
		fmt.Println("Кратно 5")
	default:
		fmt.Printf("Остаток от деления на 5: %d", j)
	}

	// EXAMPLE FALLTHROUGH
	a := -100
	switch {
	case a > 0:
		if a%2 == 0 {
			break
		}
		fmt.Println("Odd positive value received")
	case a < 0:
		fmt.Println("Negative value received")
		fallthrough // can use only in last block
	default:
		fmt.Println("Default value handling")
	}
}

func dailyWish(dayOfWeek string) (string, error) {
	switch dayOfWeek {

	case "monday":
		return "Have a good start of week", nil

	case "tuesday":
		return "Have a good tuesday", nil

	case "wednesday":
		return "Have a good wednesday", nil

	case "thursday":
		return "Have a good thursday", nil

	case "friday":
		return "Have a good friday", nil

	case "saturday":
		return "Have a good saturday", nil

	case "sanday":
		return "Have a good end of week", nil

	default:
		return "There isn't such day of week", errors.New("invalid day of the week")
	}
}
