package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	message, isEnter, err := enterTheClub(20)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message, isEnter)

	// INITIALIZATION IN IF BLOCK
	a := 0.10000001                        // float64
	if b := float32(a); b > float32(0.1) { // initialization and main condition
		fmt.Println("Var a is GT float32(0.1)")
	}
}

func enterTheClub(age int) (string, bool, error) {
	if age >= 18 && age <= 148 {
		return "Welcome to the club", true, nil
	} else if age >= 149 && age <= 150 {
		return "Welcome the oldest person in the world", true, nil
	} else if age >= 151 {
		return "Sorry, I don't believe that you are real person", false, errors.New("don't lie me, who are you?")
	}
	return "Sorry, guy. You have to be an adult", false, errors.New("sorry guy, you are too young")
}
