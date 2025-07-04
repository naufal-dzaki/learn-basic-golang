package main

import "fmt"

func main() {
	counter := 0
	for counter < 10 {
		fmt.Print(" ", counter)
		counter++
	}
	fmt.Println("")

	// for with statement
	fmt.Println("Using for with statement")
	for i := 0; i < 10; i++ {
		fmt.Print(" ", i)
	}
	fmt.Println("")

	fmt.Println("Using for to iterate over a slice")
	names := []string{"Muhammad", "Naufal", "Dzaki", "Adani", "Eko", "Kurniawan"}
	for i := 0; i < len(names); i++ {
		fmt.Print(" ", names[i])
	}

	/*
		for range loop
		- for range loop is used to iterate over elements in a slice, array, map, or string.
		- it returns the index and value of each element.

		for index, variableAlias := range collection {
			fmt.Print(" ", index, variableAlias)
		}
	*/
	fmt.Println("\nUsing for range to iterate over a slice")
	for index, name := range names {
		fmt.Print(" ", index, name)
	}

	// you can use _ to replace the index if you don't need it
	fmt.Println("\nUsing for range to iterate over a slice without index")
	for _, name := range names {
		fmt.Print(" ", name)
	}

	/*
		break
		- break statement is used to exit the loop immediately.
		- it can be used with for, switch, or select statements.
	*/
	fmt.Println("\n\nUsing break to exit the loop")
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Breaking the loop at index", i)
			break // exit the loop when i is 5
		}
		fmt.Println(" ", i)
	}

	/*
		continue
		- continue statement is used to skip the current iteration and move to the next iteration of the loop.
		- it can be used with for, switch, or select statements.
		- it is often used to skip certain conditions in the loop.
	*/
	fmt.Println("\n\nUsing continue to skip the current iteration")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // skip the iteration when i is even
		}
		fmt.Println(i)
	}

}
