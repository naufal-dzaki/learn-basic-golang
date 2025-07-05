package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("User", name, "is blocked from registration.")
	} else {
		fmt.Println("User", name, "registered successfully.")
	}
}

func main() {
	/*
		anonymous function
		- an anonymous function is a function that is defined without a name
		- you can use anonymous functions to create function literals
		- anonymous functions are often used as arguments to higher-order functions
	*/

	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("naufal", blacklist)

	registerUser("anjing", func(name string) bool {
		return name == "anjing" // inline anonymous function
	})
}
