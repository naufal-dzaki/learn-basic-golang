package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Surabaya", "Jawa Timur", "Indonesia"}
	address2 := address1 // address1 is copied into address2.

	// When address2.City is changed, address1.City stays the same — because they are two different copies.
	address2.City = "Gresik"

	fmt.Println(address1) // stay
	fmt.Println(address2) // changed

	/*
		pointer
		 What is a Pointer?
		 - A pointer is a variable that stores the memory address of another variable.
		 - It allows you to modify the original value without copying it — this is called pass by reference.

		 Why Use Pointers?
		 - To avoid copying large data structures.
		 - To allow functions or variables to modify the original data.
		 - Useful in performance-sensitive and shared-data cases.
	*/
	address3 := Address{"Sidoarjo", "Jawa Timur", "Indonesia"}
	address4 := &address3    // address4 is a pointer to address3 (using &).
	address4.City = "Gresik" // Changing address4.City also changes address3.City — because both point to the same memory.

	fmt.Println(address3) // changed
	fmt.Println(address4) // changed

	address5 := Address{"Malang", "Jawa Timur", "Indonesia"}
	var address6 *Address = &address5 // Same as above, just using a more explicit pointer declaration.
	address6.City = "Surabaya"        // address6 points to address5, so changes affect both.

	fmt.Println(address5)
	fmt.Println(address6)
}
