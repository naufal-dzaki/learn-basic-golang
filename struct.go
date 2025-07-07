package main

import "fmt"

/*
- struct is a composite data type that groups together zero or more fields (variables) under a single name.
- It is similar to an object in JavaScript, but without methods or behavior attached by default.
- Each field has a name and a type, and all fields are stored together in memory.
- Structs are used to model real-world data or represent entities in an application (e.g., User, Product, Order).
- Structs are useful for organizing related data, making code more structured and readable.
- You can create custom types based on structs, which makes your program more expressive and type-safe.
- Structs can be embedded (one struct inside another) to simulate inheritance or reuse common fields.
- Struct type names and field names typically use PascalCase (each word capitalized, no underscores), like User, ProductID, or FirstName.
- In Go, PascalCase (capitalized) identifiers are exported, meaning they are accessible from other packages.
- If you use camelCase (lowercase first letter) for struct fields, they will be unexported and only accessible within the same package.
*/
type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var customer1 Customer
	customer1.Name = "Muhammad Naufal"
	customer1.Address = "Jl. Raya No. 123"
	customer1.Age = 20
	fmt.Println("Customer 1:", customer1)
	fmt.Println("Customer 1 Name:", customer1.Name)

	// Creating a new Customer instance using struct literal
	// This is a more concise way to create and initialize a struct.
	customer2 := Customer{
		Name:    "Naufal Dzaki",
		Address: "Jl. Kebon Jeruk No. 456",
		Age:     21,
	}
	fmt.Println("Customer 2:", customer2)

	// Using a shorthand syntax to create a Customer instance
	// This is a shorthand way to create a struct without explicitly naming the fields.
	// The order of fields must match the struct definition.
	customer3 := Customer{"Dzaki Adani", "Jl. Cempaka No. 789", 22}
	fmt.Println("Customer 3:", customer3)
}
