package main

import "fmt"

/*
	recursive function
	- a recursive function is a function that calls itself
	- this is useful for solving problems that can be broken down into smaller subproblems
*/

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(5)) // Output: 120
}
