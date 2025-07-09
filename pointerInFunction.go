package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// parameter copy the value
// func ChangeAddressToIndonesia(address Address) {
// 	address.Country = "Indonesia"
// } // address never change

/*
	Pointer Function
	- By default, function parameters in Go are passed by value, meaning a copy of the data is sent to the function.
	- If you change the data inside the function, it won't affect the original variable.
	- This behavior ensures data safety — the original value stays unchanged.
	- However, if you want a function to modify the original value, you can use pointers.
	- A pointer parameter allows the function to directly reference and modify the actual data.
	- To define a pointer parameter, use the * operator before the type.
	- When calling the function, pass the memory address using &.
*/

// This function accepts a pointer to Address
// It modifies the actual Address by setting the Country field
func ChangeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{"Surabaya", "Jawa Timur", ""}
	ChangeAddressToIndonesia(&address) // Pass the memory address (reference) to the function

	fmt.Println(address) // // The original address is modified
}
