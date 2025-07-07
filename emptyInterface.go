package main

import "fmt"

/*
Empty Interface
- An empty interface is a special type in Go that can hold values of any type.
- It is defined as `interface{}` or `any` and does not specify any methods.
- Since all types implement at least zero methods, any value can be assigned to an empty interface.
- It is often used for generic programming, allowing functions to accept any type of argument.
*/

func Ops(value any) any {
	switch v := value.(type) {
	case int:
		return v * 2 // Example operation for int
	case string:
		return fmt.Sprintf("String: %s", v)
	case bool:
		return fmt.Sprintf("Boolean: %t", v)
	default:
		return "Unknown type"
	}
}

func main() {
	fmt.Println(Ops(10))
	fmt.Println(Ops("Hello"))
	fmt.Println(Ops(true))
	fmt.Println(Ops(3.14))
}
