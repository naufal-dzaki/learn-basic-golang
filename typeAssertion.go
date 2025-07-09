package main

import "fmt"

func random() any {
	return "hello"
}

func main() {
	var result any = random()

	/*
		Type Assertion
		- Type assertion is a way to retrieve the dynamic type of an interface variable in Go.
		- It allows you to check if the interface holds a specific type and to extract that value.
		- The syntax for type assertion is `value, ok := interfaceVariable.(Type)`, where `ok` is a boolean that indicates whether the assertion was successful.
		- If the assertion is successful, `value` will hold the value of the asserted type; if it fails, `value` will be the zero value of that type and `ok` will be false.
		- If you are sure about the type and want to avoid the boolean check, you can use `value := interfaceVariable.(Type)`, but this will panic if the assertion fails.
		- Type assertions are commonly used when working with interfaces to handle different types dynamically.
	*/
	// var resultString string = result.(string) // Type assertion to convert 'any' to 'string'
	// fmt.Println(resultString) // Output: hello

	// var resultInt int = result.(int) // This will panic if 'result' is not an int
	// fmt.Println(resultInt) // Output: 0 (if 'result' is not an int

	// Using type assertion with a check
	resultString, ok := result.(string)
	if ok {
		fmt.Println(resultString) // Output: hello
	} else {
		fmt.Println("Type assertion failed")
	}

	switch value := result.(type) {
	case string:
		fmt.Println("String value:", value) // Output: String value: hello
	case int:
		fmt.Println("Integer value:", value)
	default:
		fmt.Println("Unknown type")
	}

}
