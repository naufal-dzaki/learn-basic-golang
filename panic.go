package main

import "fmt"

/*
	panic function
	- a panic function is used to stop the normal execution of a program
	- it is typically used for unrecoverable errors or unexpected conditions
	- panic is usually triggered when something goes seriously wrong during runtime
	- when a panic is triggered, the program starts to unwind and eventually stops
	- however, any deferred functions will still be executed before the program fully stops
*/

func endApp() {
	fmt.Println("Application ended")
}

func runApp(error bool) {
	defer endApp() // this function will be executed even if panic occurs
	fmt.Println("Application is running")

	if error {
		panic("An error occurred") // trigger a panic
	}

	fmt.Println("Application completed successfully")

}

func main() {
	runApp(true) // this will trigger a panic
	fmt.Println("This line will not be executed due to panic")
}
