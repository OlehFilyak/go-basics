package main

import (
	"fmt"
)

type Age int

func (age Age) isAdult() bool {
	return age >= 18
}

type User struct {
	name   string
	age    Age
	sex    string
	weight int
	height int
}

func NewUser(name, sex string, age, weight, height int) User {
	return User{
		name:   name,
		sex:    sex,
		age:    Age(age),
		weight: weight,
		height: height,
	}
}

func (user *User) setName(name string) {
	user.name = name
}

func (user User) getName() string {
	return user.name
}

// func (user User) printUserInfo() { //value
// 	fmt.Println(user.name, user.age, user.sex, user.height)
// }

// func (user *User) printUserInfo() { //pointer
// 	fmt.Println(user.name, user.age, user.sex, user.height)
// }

// type DumbDatabase struct {
// 	m map[string]string
// }

// func NewDumbDatabase() *DumbDatabase {
// 	return &DumbDatabase{
// 		m: make(map[string]string),
// 	}
// }

func main() {
	// user := struct {
	// 	name   string
	// 	age    int
	// 	sex    string
	// 	weight int
	// 	height int
	// }{"Boryviter", 55, "Male", 75, 177}

	// fmt.Println(user)
	// fmt.Printf("%+v\n", user)

	user1 := NewUser("Ivan", "Male", 23, 77, 199)
	user2 := User{"Petro", 25, "Male", 79, 184}

	// fmt.Println(user1.weight)
	// fmt.Printf("%+v\n", user1)
	// fmt.Printf("%+v\n", user2)

	// printUserInfo(user1)
	// printUserInfo(user2)

	// user1.printUserInfo()
	// user2.printUserInfo()

	fmt.Println(user1.age.isAdult())
	fmt.Println(user1.getName())
	fmt.Println(user2.getName())
	user2.setName("Ioan")
	fmt.Println(user2.getName())
}

// func printUserInfo(user User) {
// 	fmt.Println(user.name, user.age, user.sex, user.height)
// }
