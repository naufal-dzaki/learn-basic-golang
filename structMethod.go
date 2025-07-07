package main

import "fmt"

/*
struct method
- A struct method is a function that is associated with a specific struct type.
- It allows you to define behavior or operations that can be performed on instances of that struct.
- Struct methods are defined with a receiver, which is a special parameter that specifies the struct type the method belongs to.
- The receiver is defined before the method name and is similar to the `this` keyword in other programming languages.
- Struct methods can access and modify the fields of the struct they are associated with.
- They can also return values, just like regular functions.
*/

type Customer struct {
	Name, Address string
	Age           int
}

// sayHello is a method of the Customer struct.
// It takes a name as an argument and prints a greeting message.
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	customer := Customer{"Naufal Dzaki", "Jl. Kebon Jeruk No. 456", 21}
	customer.sayHello("Adani")
}
