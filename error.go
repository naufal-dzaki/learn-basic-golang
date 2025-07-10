package main

/*
Error
- Go handles errors explicitly using the error type.
- The errors.New() function from the errors package creates a new error message.
- A common Go convention is to return a tuple: (result, error) from a function.
- The caller of the function must check if the error is nil before using the result.
- This approach improves clarity and control by forcing developers to handle errors explicitly, unlike exceptions in other languages.
*/

import (
	"errors"
	"fmt"
)

func divide(value, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("division by zero is not allowed")
	} else {
		return value / divisor, nil // If no error, returns the result and nil.
	}
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result) // Output: Result: 5
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err) // Output: Error: division by zero is not allowed
	} else {
		fmt.Println("Result:", result)
	}
}
