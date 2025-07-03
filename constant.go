package main

import "fmt"

func main() {
	/*
		const declarations can be unused without causing a compile-time error.
		However, declared variables (var) must be used, or the compiler will raise an error.
	*/
	const firstName string = "Muhammad Naufal"
	const lastName = "Dzaki Adani"

	// This will cause an error because firstName is a constant and cannot be reassigned
	// firstName = "Naufal"

	fmt.Println(firstName)

	// multiple constants can be declared at once using the const keyword and parentheses
	const (
		location  = "Indonesia"
		isMarried = false
	)

	fmt.Println(location)
	fmt.Println(isMarried)
}
