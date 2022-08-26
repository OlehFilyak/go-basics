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
