package main

import "fmt"

/*
	defer
	- defer is a keyword in Go that allows you to schedule a function call to be run after the surrounding function completes
	- deferred function calls are executed in LIFO (last-in, first-out) order
	- you can use defer to ensure that resources are released, files are closed, or other cleanup tasks are performed
	- defer is often used for error handling and resource management
	- deferred functions will still run even if an error occurs in the main function
*/

func logging() {
	fmt.Println("Logging function executed")
}

func runApplication() {
	defer logging() // this function will be executed after runApplication completes
	fmt.Println("Running application")
	// you can have multiple defer statements
	defer fmt.Println("This will be printed last")
	fmt.Println("Application is running")
}

func main() {
	runApplication()
}
