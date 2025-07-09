package main

import "fmt"

type Man struct {
	Name string
}

// func (man Man) Married() {
// 	man.Name = "Mr. " + man.Name
// } // This does NOT change the original Name because it's a copy

/*
	Pointer In Method
	- A method is a function that is associated with a type (like a struct).
	- By default, methods in Go receive a copy of the struct.
	- If you want to modify the original struct, you must use a pointer receiver.
	- A pointer receiver allows the method to access and modify the actual data, not just a copy.
	- This is useful for methods that mutate state (change the struct's fields).
*/

// Method with pointer receiver (reference to struct)
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
} // This DOES change the original Name because it's a pointer

func main() {
	eko := Man{"Naufal"}
	eko.Married()

	fmt.Println(eko)
}
