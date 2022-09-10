package main

import (
	"fmt"
)

func main() {
	var users1 map[string]int
	fmt.Println(users1)
	users1 = make(map[string]int) // For an object of reference type to be usable
	users1["Vasya"] = 22          // it must first be created (initialized) with the built-in make() function
	fmt.Println(users1)
	fmt.Println(len(users1))

	// var m map[string]string
	// if m != nil { // if you do not check this condition,
	// 	m["foo"] = "bar" // you can get panic here
	// }
	// fmt.Println(m)

	// var MyMap map[[]byte]string // error

	m := map[int]string{99: "first"}
	v, ok := m[99]
	// v, ok = m[k]
	// v, ok := m[k]
	// var v, ok = m[k]
	// addr := &m[k] // error
	fmt.Println(m)
	fmt.Println(v, ok)
	delete(m, 1)
	v, ok = m[1]
	fmt.Println(v, ok)

	t := make(map[string]string)
	t["foo"] = "bar"
	t["bazz"] = "yup"
	for k, v := range t {
		// k буде перебирати ключі,
		// v — значення відповідні до ключів
		// v = "here key " + k   // Не спрацює для модифікації таблиці.
		t[k] = "here key " + k // Спрацює для модифікації таблиці.
		fmt.Printf("Ключ %v, має значення %v \n", k, v)
	}

	// sentence
	sentence := "Πολύ λίγα πράγματα συμβαίνουν στο σωστό χρόνο, και τα υπόλοιπα δεν συμβαίνουν καθόλου"
	// initialize the map
	// the keys are the signs, and the values are the number of occurrences
	frequency := make(map[rune]int)
	// let's walk through the signs in the sentence
	for _, v := range sentence {
		frequency[v]++ //increase frequency
	}
	// print
	for k, v := range frequency {
		fmt.Printf("The sign %c occurs %d times \n", k, v)
	}

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
