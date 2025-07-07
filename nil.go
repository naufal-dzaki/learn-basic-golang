package main

import "fmt"

/*
nil
- `nil` is a predeclared identifier in Go that represents the zero value for pointers, interfaces, maps, slices, channels, and function types.
- It is used to indicate the absence of a value or a null reference.
- When a variable of one of these types is declared but not initialized, it is automatically set to `nil`.
- You can check if a variable is `nil` using the comparison operator `==`.
*/

func NewMap(name string) map[string]string {
	if name == "" {
		return nil // Return nil if the name is empty
	}
	return map[string]string{"name": name} // Return a new map with the provided name
}

func main() {
	data := NewMap("")
	if data == nil {
		fmt.Println("Data is nil") // This will be printed because NewMap returns nil for an empty name
	} else {
		fmt.Println("Data:", data) // This won't be executed in this case
	}

	data2 := NewMap("Naufal Dzaki")
	if data2 == nil {
		fmt.Println("Data2 is nil") // This won't be printed because data2 is not nil
	} else {
		fmt.Println("Data2:", data2) // This will print the map with the name "Naufal Dzaki"
	}
}
