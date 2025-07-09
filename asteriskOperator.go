package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Surabaya", "Jawa Timur", "Indonesia"}
	address2 := &address1 // address2 := &address1 makes address2 point to address1.

	//  you reassign address2 to point to a new Address, so address1 stays the same. address2 now points to the new address (Gresik).
	address2 = &Address{"Gresik", "Jawa Timur", "Indonesia"}
	fmt.Println(address1) // stay
	fmt.Println(address2) // changed

	/*
		Asterisk Operator
		 What is the Asterisk (*) Operator?
		 In Go, the * operator is used in two main ways:
		 1. Declaration (Pointer Type):
			When used in a type declaration (*Type), it defines a pointer to that type.
			var ptr *int // pointer to an int

		 2. Dereferencing (Access or Modify Value):
		 When used with a pointer variable (*ptr), it means "access or set the value at that memory address."
		 *ptr = 10 // set the value the pointer points to
		 fmt.Println(*ptr) // get the value the pointer points to
	*/
	address3 := Address{"Surabaya", "Jawa Timur", "Indonesia"}
	address4 := &address3 // address4 points to address3

	// replace the value stored at the memory address of address3
	// So both address3 and address4 now point to the same new value (Gresik).
	*address4 = Address{"Gresik", "Jawa Timur", "Indonesia"}
	fmt.Println(address3) // changed
	fmt.Println(address4) // changed
}
