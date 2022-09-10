package main

import (
	"fmt"
)

func main() {
	// INITIALIZATION 1
	// p := Person{}
	// or
	// var p Person

	// INITIALIZATION 2
	// date := time.Date(2000, 12, 1, 0, 0, 0, 0, time.UTC)
	// p := Person{ "Иван", "ivan@yandex.ru", date }

	// INITIALIZATION 3
	// p := Person{Name: "Иван", Email: "ivan@yandex.ru"}

	// func NewPerson(name, email string, dobYear, dobMonth, dobDate int) Person {
	// 	return Person{
	// 		Name:        name,
	// 		Email:       email,
	// 		dateOfBirth: time.Date(dobYear, time.Month(dobMonth), dobDate, 0, 0, 0, 0, time.UTC),
	// }
	// }
	fmt.Println()
}
