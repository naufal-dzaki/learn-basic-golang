package main

import "fmt"

func main() {
	/*
		closure
		- a closure is a function that captures the variables from its surrounding scope
		- this allows the function to access and modify those variables even after the surrounding scope has finished executing
		- closures are often used to create functions with state, such as counters or accumulators
		- closures can be used to create private variables that are not accessible from outside the function
		- use closures wisely — overusing them can make your code harder to read and maintain
	*/

	counter := 0

	increment := func() {
		fmt.Println("Incrementing counter")
		// this function captures the counter variable from the surrounding scope
		counter++ // increment the counter variable
	}

	increment()
	fmt.Println("Counter value:", counter) // prints: Counter value: 1
	increment()
	increment()
	fmt.Println("Counter value:", counter) // prints: Counter value: 2
}
