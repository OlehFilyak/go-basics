package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
}

type Square struct {
	sideLength float32
}

func (s Square) Area() float32 {
	return s.sideLength * s.sideLength
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func main() {
	square := Square{5}
	circle := Circle{8}

	printShapeArea(square)
	printShapeArea(circle)

	printInterface(square)
	printInterface(circle)
	printInterface("qwe")
	printInterface(22)
	printInterface(true)
}

func printShapeArea(shape Shape) {
	fmt.Println(shape.Area())

	// fmt.Println("fwerfsw", 333, 4242.423, true)
}

func printInterface(i interface{}) {
	// fmt.Printf("%+v\n", i)

	// switch value := i.(type) {
	// case int:
	// 	fmt.Println("int", value)
	// case bool:
	// 	fmt.Println("bool", value)
	// case string:
	// 	fmt.Println("string", value)
	// case float32:
	// 	fmt.Println("float32", value)
	// default:
	// 	fmt.Println("another type", value)
	// }

	str, ok := i.(string)
	if !ok {
		fmt.Println("interface is not string")
		return
	}

	fmt.Println(len(str))
}
