package main

import "fmt"

func getGoodBye() string {
	return "Goodbye"
}

func main() {
	/*
		function as value
		- in Go, functions are first-class citizens, which means they can be assigned to variables
		- you can pass functions as arguments to other functions
		- you can return functions from other functions
	*/

	goodBye := getGoodBye // assigning a function to a variable
	// you can assign a function to a variable without calling it
	fmt.Println(goodBye()) // calling the function using the variable

	// you can also assign a function to a variable and call it later
	// this is useful when you want to pass a function as an argument to another function
	// assigning a function to a variable
	sayHello := func(firstName, lastName string) {
		fmt.Println("Hello", firstName, lastName)
	}

	// calling the function
	sayHello("Naufal", "Dzaki")
}
