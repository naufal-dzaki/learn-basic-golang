package main

import "fmt"

func sayHello(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func getHello(firstName string, lastName string) string {
	return "Hello " + firstName + " " + lastName
}

func getFullName() (string, string) { // this function returns two values, a string and a string
	// you can use this function to return multiple values
	return "Naufal", "Dzaki"
}

func main() {
	sayHello("Naufal", "Dzaki")              // calling the sayHello function
	fmt.Println(getHello("Naufal", "Dzaki")) // calling the getHello function

	firstName, lastName := getFullName() // calling the getFullName function
	fmt.Println(firstName, lastName)     // printing the returned values

	/*
		when you have a function that returns multiple values,
		you can use the blank identifier _ to ignore some  return value
	*/
	firstName2, _ := getFullName() // calling the getFullName function, ignoring the second return value
	fmt.Println(firstName2)        // printing the first return value
}
