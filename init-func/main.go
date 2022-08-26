package main

import (
	"fmt"
)

var msg string

func main() {
	msg = "from func main"
	fmt.Println(msg)
}

func init() {
	msg = "from func init"
	fmt.Println(msg)
}
