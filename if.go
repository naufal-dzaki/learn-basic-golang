package main

import "fmt"

func main() {
	name := "naufal"

	if name == "naufal" {
		fmt.Println("Hello Naufal")
	} else if name == "dzaki" {
		fmt.Println("Hello Dzaki")
	} else {
		fmt.Println("Hello Guest")
	}

	/*
		if short statement
		- if statement can have a short statement before the condition.
	*/

	if length := len(name); length > 5 {
		fmt.Println("Name is longer than 5 characters")
	} else {
		fmt.Println("Name is shorter than or equal to 5 characters")
	}
}
