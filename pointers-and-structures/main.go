package main

import (
	"fmt"
)

func main() {
	// implicit dereferencing when accessing structure fields.
	// type A struct {
	// 	IntField int
	// }

	// p := &A{}
	// p.IntField = 42 // Instead of (*p).IntField = 42

	// When to use pointers
	incrementCopy := func(i int) {
		i++
	}

	increment := func(i *int) {
		(*i)++
	}

	i := 42

	incrementCopy(i)
	fmt.Println(i) // 42

	increment(&i)
	fmt.Println(i) // 43

	// 	// EXAMPLE
	// 	type Person struct {
	// 		Name        string
	// 		Age         int
	// 		lastVisited time.Time
	// 	}

	// 	func UpdatePersonWithLastVisisted (p *Person )  {
	//   p.lastVisited = time.Now()
	// 	}

	// 	p := P{
	// 		Name: "Alex",
	// 		Age: 25,
	// 		lastVisited: time.Time{}
	// 	}

	// UpdatePersonWithLastVisisted(&p)

	// TASK
	a := 1
	p := &a //1 //3 //5
	b := &p //1 //3 //5

	*p = 3  //3 //5
	**b = 5 //5

	a += 4 + *p + **b // a = a + 4 + 5 + 5 // 19

	fmt.Println(a, p, b)
	fmt.Printf("%d", *p) // ?

}
