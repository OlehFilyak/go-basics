package main

import "fmt"

func main() {
	a, b, c, d := 1, 2, 3, true
	// a, b = b, a // поміняли змінні однією операцією
	a, _, c, d = 99, 6, 100, false // пропуск перезапису змінної

	fmt.Println(a, b, c, d)
}
