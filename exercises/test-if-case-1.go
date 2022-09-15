package main

import "fmt"

// По данному трехзначному числу определите, все ли его цифры различны.

// Формат входных данных
// На вход подается одно натуральное трехзначное число.

// Формат выходных данных
// Выведите "YES", если все цифры числа различны, в противном случае - "NO".

// Sample Input 1:
// 237
// Sample Output 1:
// YES
// Sample Input 2:
// 117
// Sample Output 2:
// NO

func main() {
	var a int
	fmt.Scan(&a)

	b := a % 10
	fmt.Println(b)
	c := (a - b) / 10 % 10
	fmt.Println(c)
	d := (a - c) / 10 % 10
	fmt.Println(d)

	if b == c && c == d && b == d {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}
