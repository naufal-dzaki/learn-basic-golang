package main

import "fmt"

func main() {
	/*
		Creating a slice from an array
		- array[start:end] creates a slice from the array from index `start` to `end-1`.
		- array[start:] creates a slice from the array from index `start` to the end of the array.
		- array[:end] creates a slice from the beginning of the array to index `end-1`.
		- array[:] creates a slice from the beginning of the array to the end of the array.
	*/

	names := [...]string{"Muhammad", "Naufal", "Dzaki", "Adani", "Eko", "Kurniawan"}
	slice := names[1:3]

	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice)

	fmt.Println(len(slice)) // len() function returns the length of the slice
	/*
		cap() function returns the capacity of the slice, which is the number of elements it can hold.
		cap count for slice is the number of elements from the start index of the slice to the end of the original array.
		In this case, the slice starts at index 1 and goes to the end of the
	*/
	fmt.Println(cap(slice))

	var slice2 []string = names[:] // if you create a slice, you can declare without specifying the size, it differs from an array
	fmt.Println(slice2)

	/*
		append(slice, data) function adds an element to the end of the slice.
		It returns a new slice with the added element.
		If the slice has enough capacity, it will not create a new array.
		If the slice does not have enough capacity, it will create a new array with double the capacity.
	*/
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice := days[5:]
	// if you change the value of an element in a slice, it will also change the value in the original array
	daySlice[0] = "Sabtu Baru"
	daySlice[1] = "Minggu Baru"
	fmt.Println(days)
	fmt.Println(daySlice)

	daySlice2 := append(daySlice, "Hari Baru") // adding a new element to the slice
	daySlice2[0] = "Whops"                     // this will not change the original array, because daySlice2 is a new slice with a new array
	fmt.Println(daySlice2)
	fmt.Println(daySlice)
	fmt.Println(days)

	/*
		make([]TypeData, length, capacity) function creates a new slice with the specified length and capacity.
		If the capacity is not specified, it will be the same as the length.
	*/
	newSlice := make([]string, 2, 5) // creates a new slice with length 2 and capacity 5
	newSlice[0] = "Naufal"
	newSlice[1] = "Dzaki"
	// newSlice[2] = "Adani" // this will cause a runtime error: index out of range, if you want to add more elements, you can use append() function
	newSlice = append(newSlice, "Adani") // now we can add more elements to the slice

	fmt.Println(newSlice)      // [Naufal Dzaki Adani]
	fmt.Println(len(newSlice)) // 3
	fmt.Println(cap(newSlice)) // 5

	/*
		copy(destination, source) function copies the elements from the source slice to the destination slice.
	*/
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice)) // create a new slice with the same length as fromSlice
	copy(toSlice, fromSlice)                                  // copy the elements from fromSlice to toSlice
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	thisIsArray := [...]int{1, 2, 3, 4, 5} // this is an array, not a slice
	thisIsSlice := []int{1, 2, 3, 4, 5}    // this is a slice, it can be declared without specifying the size
	fmt.Println(thisIsArray)
	fmt.Println(thisIsSlice)
}
