package main

import (
	"fmt"
)

/*
	recover function
	- recover is a built-in function that allows you to regain control of a panicking goroutine
	- it can be used to handle panics and prevent the program from crashing
	- recover is used to catch the panic value inside a deferred function
	- when recover is called, the panic stops and the program continues to run normally
*/

// incorrect recover usage
// func endApp() {
// 	fmt.Println("Application ended")
// }

// func runApp(error bool) {
// 	defer endApp()
// 	fmt.Println("Application is running")

// 	if error {
// 		panic("An error occurred")
// 	}

// 	message := recover() // this will not catch the panic because it's not deferred
// 	if message != nil {
// 		fmt.Println("Recovered from panic:", message)
// 	}

// 	fmt.Println("Application completed successfully")
// }

// correct recover usage
func endApp() {
	fmt.Println("Application ended")
	message := recover() // recover should be called inside a deferred function
	if message != nil {
		fmt.Println("Recovered from panic:", message)
	}
}

func runApp(error bool) {
	defer endApp()
	fmt.Println("Application is running")

	if error {
		panic("An error occurred")
	}

	fmt.Println("Application completed successfully")
}

func main() {
	runApp(true) // this will trigger a panic
	fmt.Println("This line will executed after recover if panic occurs")
}
