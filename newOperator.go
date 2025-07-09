package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 *Address = &Address{} // &Address{} creates a pointer to a new empty Address struct.
	address2 := address1               // address1 and address2 both point to the same memory.

	address2.Country = "Indonesia" // Changing address2.Country also changes address1.Country.

	fmt.Println(address1)
	fmt.Println(address2)

	/*
		new operator
		 What is the new Operator?
		 - The new operator in Go is used to allocate memory for a value of a given type and returns a pointer to it.
		 - The allocated value is zero-initialized (all fields set to their zero values).
	*/

	var address3 *Address = new(Address) // new(Address) also returns a pointer to a zero-valued Address struct.
	address4 := address3                 // address3 and address4 point to the same memory.

	address4.Country = "Indonesia" // So changing address4.Country affects address3 as well.

	fmt.Println(address3)
	fmt.Println(address4)

}
