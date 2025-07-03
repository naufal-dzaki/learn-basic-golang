package main

import "fmt"

func main() {
	/*
		all declared variables must be used. If a variable is declared but not used, the compiler will throw an error.
	*/
	var name string               // if you want to declare a variable without a value, you must specify the type
	var username = "naufal-dzaki" // when declaring a variable with a value, you can omit the type
	/*
		when declaring a variable with a value,
		you can omit the type and the keyword var,
		but you must use the short declaration operator :=
	*/
	igname := "naufal.dza"

	name = "Muhammad Naufal"
	fmt.Println(name)

	name = "Dzaki Adani"
	fmt.Println(name)

	fmt.Println(username)

	fmt.Println(igname)

	var (
		firstName = "Muhammad Naufal"
		lastName  = "Dzaki Adani"
		fullName  = firstName + " " + lastName
	) // you can declare multiple variables at once using var and parentheses

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(fullName)
}
