package main

import "fmt"

/*
	function as parameter
	- in Go, you can pass functions as parameters to other functions
	- this allows you to create higher-order functions that can operate on other functions
*/
// func sayHelloWithFilter(name string, filter func(string) string) {
// 	fmt.Println("Hello", filter(name))
// }

/*
	function type declaration
	- you can declare a function type to make the code more readable
	- this is useful when you want to pass a function as a parameter to another function
*/
type NameFilter func(string) string

func sayHelloWithFilter(name string, filter NameFilter) {
	fmt.Println("Hello", filter(name))
}

func filterName(name string) string {
	// this function is used to filter the name
	// for example, if the name is "anjing", we will replace it with "*****"
	if name == "anjing" {
		return "*****" // replace "anjing" with "*****"
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("anjing", filterName)

	filter := filterName
	sayHelloWithFilter("Naufal", filter)
}
