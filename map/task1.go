package main

import (
	"fmt"
)

// Now try it yourself. Make a map to store the price list in grivnas:
// bread - 50,
// milk - 100,
// butter - 200,
// sausage - 500,
// salt - 20,
// cucumbers 200,
// cheese 600,
// ham 700,
// burger 900,
// tomatoes 250,
// fish - 300,
// Jamon - 1500.
// Derive a list of delicacies - products more expensive than 500 grivnas.
// The order is given by slice []string{"bread", "baked ham", "cheese", "cucumbers"}. Calculate the cost of the order.

func main() {
	groceries := map[string]int{
		"bread":     50,
		"milk":      100,
		"butter":    200,
		"sausage":   500,
		"cucumbers": 200,
		"cheese":    600,
		"ham":       700,
		"burger":    900,
		"tomatoes":  250,
		"fish":      300,
		"Jamon":     1500,
	}

	order := make(map[string]int)
	orderPrice := 0

	groceriesExpensive := make(map[string]int)

	for k, v := range groceries {

		if v > 500 {
			groceriesExpensive[k] = v
		}

		if k == "bread" || k == "ham" || k == "cucumbers" || k == "cheese" {
			order[k] = v
		}

		fmt.Printf("Ключ %v, має значення %v \n", k, v)

	}
	fmt.Println(groceriesExpensive)

	for _, v := range order {
		orderPrice += v
	}
	fmt.Println(order)
	fmt.Println(orderPrice)
	// fmt.Println(groceries)
}
