package main

import "fmt"

func main() {
	/*
		In this code, an int32 value is converted to int64, and then to int16.
		Type conversion in Go is explicit using syntax like int16(value).
		Conversion between different numeric types in Go requires explicit type conversion.
	*/
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	/*
		When converting between numeric types, make sure the value fits within the target type's capacity.
		If it exceeds the limit, the value wraps around (overflows), leading to unexpected results.
	*/
	var nilai16 int16 = int16(nilai64)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "Naufal Dzaki"
	var firstWord = name[0] // firstWord is of type byte, which is an alias for uint8
	// To convert byte to string, you can use the string() function
	var firstWordString = string(firstWord)

	fmt.Println(name)
	fmt.Println(firstWord)       // Output: 78 (ASCII value of 'N')
	fmt.Println(firstWordString) // Output: "N"
}
