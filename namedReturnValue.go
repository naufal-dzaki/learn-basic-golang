package main

import "fmt"

func getCompleteName() (firstName, lastName string) { // this function uses named return values
	// you can use named return values to make the code more readable
	firstName = "Naufal"
	lastName = "Dzaki"
	return // you can use named return values, so you don't need to specify the return values
}

func main() {
	/*
		you can also use named return values, which allows you to return values with names
		you can use the names in the function body, and they will be returned automatically
		you can also use the names in the return statement, but it's not necessary
	*/
	firstName, lastName := getCompleteName() // calling the getCompleteName function
	fmt.Println(firstName, lastName)         // printing the returned values
}
