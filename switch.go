package main

import "fmt"

func main() {
	name := "naufal"

	switch name {
	case "naufal":
		fmt.Println("Hello Naufal")
	case "dzaki":
		fmt.Println("Hello Dzaki")
	default:
		fmt.Println("Hello Guest")
	}

	/*
		switch case can have a short statement before the condition.
	*/

	switch length := len(name); {
	case length > 5:
		fmt.Println("Name is longer than 5 characters")
	case length == 5:
		fmt.Println("Name is exactly 5 characters")
	default:
		fmt.Println("Name is shorter than 5 characters")
	}

}
