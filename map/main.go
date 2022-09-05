package main

import (
	"fmt"
)

func main() {
	var users1 map[string]int
	fmt.Println(users1)
	users1 = make(map[string]int)
	users1["Vasya"] = 22
	fmt.Println(users1)
	fmt.Println(len(users1))

	users := map[string]int{
		"Ivan":  33,
		"Petro": 45,
		"Slava": 99,
	}

	age, exists := users["Kostya"]
	fmt.Println(age, exists)
	if exists {
		fmt.Println("Kostya", age)
	} else {
		fmt.Println("Kostya не має в списку")
	}

	age1, exists1 := users["Ivan"]
	fmt.Println(age1, exists1)
	if exists1 {
		fmt.Println("Ivan", age1)
	} else {
		fmt.Println("Ivan не має в списку")
	}

	users["Serega"] = 41
	delete(users, "Ivan")

	for key, value := range users {
		fmt.Println(key, value)
	}
}
