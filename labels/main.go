package main

import (
	"fmt"
)

func main() {
	// LABELS BREAK, CONTINUE, GOTO

	// BREAK
outerLoopLabelBreak:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("[%d, %d]\n", i, j)
			break outerLoopLabelBreak // Here break outerLoopLabelBreak interrupts the execution of the outer loop.
		}
	}
	fmt.Println("End")

outerLoopLabelContinue:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("[%d, %d]\n", i, j)
			continue outerLoopLabelContinue // Here continue outerLoopLabelContinue causes the transition to the next iteration of the outer loop.
		}
	}
	fmt.Println("End")

	// EVEN NUMBERS
	group := 0
	for i := 0; i < 20; i++ {
		switch {
		case i%2 == 0:
			if i%10 == 0 {
				group++
				break // break относится к ближайшему switch
			}
			fmt.Printf("%02d: %d\n", group, i)
		default:
		}
	}

	// FIZZBUZZ
	// Here is a classic job for interviews in any programming language - FizzBuzz.
	// Write a program that displays numbers from 1 to 100.
	// The program should print the word Fizz instead of numbers divisible by three, and Buzz instead of numbers divisible by five.
	// If the number is a multiple of both three and five, the program should print the word FizzBuzz.
	// When you have completed the problem, click "Done" and compare your version with the suggested solution.
	for i := 1; i <= 100; i++ {
		found := false

		if i%3 == 0 {
			fmt.Printf("Fizz")
			found = true
		}
		if i%5 == 0 {
			fmt.Printf("Buzz")
			found = true
		}

		if !found {
			fmt.Println(i)
			continue
		}

		fmt.Println()
	}
}
