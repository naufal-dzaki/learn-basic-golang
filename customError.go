package main

/*
Custom Error
- In Go, error is an interface with a single method:
	type error interface {
		Error() string
	}
- To create custom errors, define a struct and implement the Error() method so that the struct satisfies the error interface.
- This allows for more descriptive and typed error handling using type assertion or switch type.
- Custom error types are useful for distinguishing between different kinds of errors (e.g., ValidationError, NotFoundError), enabling more granular control.
*/

import (
	"fmt"
)

// validationError is a custom error type that returns its message via the Error() method.
type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

// notFoundError is a custom error type that returns its message via the Error() method.
type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func saveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "validation error"}
	}

	if id != "naufal" {
		return &notFoundError{Message: "Data Not Found"}
	}

	return nil
}

func main() {
	err := saveData("", nil)
	if err != nil { // error found
		// if validationErr, ok := err.(*validationError); ok {
		// 	fmt.Println("Validation error: ", validationErr.Error())
		// } else if notFoundErr, ok := err.(*notFoundError); ok {
		// 	fmt.Println("Not found error: ", notFoundErr.Error())
		// } else {
		// 	fmt.Println("Unknown Error: ", err.Error())
		// }

		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("Validation error: ", finalError.Error())
		case *notFoundError:
			fmt.Println("Not found error: ", finalError.Error())
		default:
			fmt.Println("Unknown error", finalError.Error())
		}
	} else {
		fmt.Println("Success")
	}
}
