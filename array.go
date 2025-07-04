package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Naufal"
	names[1] = "Dzaki"
	names[2] = "Adani"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	values := [3]int{
		90,
		80,
		70, // if use multiple lines, the last line must have a comma
	}
	fmt.Println(values)

	// if you declare an array with less elements than its size, the rest will be filled with zero values
	numbers := [5]int{1, 2, 3, 4}
	fmt.Println(numbers)

	numbers[0] = 100 // you can change the value of an array element
	fmt.Println(numbers)

	fmt.Println(len(numbers)) // len() function returns the length of the array
	fmt.Println(cap(numbers)) // cap() function returns the capacity of the array, which is the number of elements it can hold

	ages := [...]int{20, 21, 22, 23, 24} // [...] means the size of the array will be determined by the number of elements
	fmt.Println(ages)

}
