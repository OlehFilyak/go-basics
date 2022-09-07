package main

import (
	"fmt"
	// "reflect"
)

func main() {
	var someString string
	var SomeInt int
	SomeInt = 2
	SomeInt *= 2 // 2*2 = 4
	SomeInt++    // SomeInt = SomeInt + 1
	SomeInt--    // SomeInt = SomeInt - 1
	var someUint uint
	var SomeUintptr uintptr
	var someFloat32 float32
	var someFloat64 float64
	var someInt8 int8
	var someInt16 int16
	var someInt32 int32
	var someInt64 int64
	var someUint8 uint8
	var someUint16 uint16
	var someUint32 uint32
	var someUint64 uint64
	var someByte byte
	var someRune rune
	var someComplex64 complex64
	var someComplex128 complex128
	var someBool bool

	//Zero values
	fmt.Println("'", someString, "'") // \n - перехід на нову строку, \r - повернення каретки, \t - табуляція, \" - подвійна лапка в середині рядків, \\ - зворотній слеш
	fmt.Println(SomeInt)              // unsigned, either 32 or 64 bits
	fmt.Println(someUint)             // signed, either 32 or 64 bits
	fmt.Println(SomeUintptr)          //unsigned integer large enough to store the uninterpreted bits of a pointer value
	fmt.Println(someFloat32)          //IEEE-754 32-bit floating-point numbers
	fmt.Println(someFloat64)          //IEEE-754 64-bit floating-point numbers
	fmt.Println(someInt8)             //signed  8-bit integers (-128 to 127)
	fmt.Println(someInt16)            //signed 16-bit integers (-32768 to 32767)
	fmt.Println(someInt32)            //signed 32-bit integers (-2147483648 to 2147483647)
	fmt.Println(someInt64)            //signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
	fmt.Println(someUint8)            //unsigned  8-bit integers (0 to 255)
	fmt.Println(someUint16)           //unsigned 16-bit integers (0 to 65535)
	fmt.Println(someUint32)           //unsigned 32-bit integers (0 to 4294967295)
	fmt.Println(someUint64)           //unsigned 64-bit integers (0 to 18446744073709551615)
	fmt.Println(someByte)             //alias for uint8
	fmt.Println(someRune)             //alias for int32
	fmt.Println(someComplex64)
	fmt.Println(someComplex128)
	fmt.Println(someBool) // true or false

	// fmt.Println(reflect.TypeOf(firstName))
	// fmt.Println(reflect.TypeOf(birthadayYear))
}
