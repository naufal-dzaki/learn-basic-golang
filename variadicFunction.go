package main

import "fmt"

// function with slice parameter
func sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

/*
	variadic function is a function that can accept a variable number of arguments in the last parameter
	you can use the ... operator to define a variadic function
*/

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	// calling the sum function with a slice
	fmt.Println(sum([]int{1, 2, 3, 4, 5})) // printing the total of the slice

	/*
		you can easily convert a slice to a variadic function
		variadic function is a function that can accept a variable number of arguments
		you can use the ... operator to define a variadic function
		you can call the function with a variable number of arguments
		or with a slice by using the ... operator
	*/

	total := sumAll(1, 2, 3, 4, 5) // calling the sumAll function with a variable number of arguments
	fmt.Println(total)             // printing the total

	// you can also call the function with a slice by using the ... operator
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	total = sumAll(slice...) // calling the sumAll function with a slice
	fmt.Println(total)       // printing the total
}
