package main

import "fmt"

func main() {
	/*
		map is a collection of key-value pairs, where each key is unique and maps to a value.
		It is different from a slice or an array, which are indexed by integers.

		Unlike arrays and slices, a map allows you to add as many entries as you want,
		as long as each key is unique. If you insert a value using an existing key,
		the previous value will be automatically replaced.

		Creating a map
		- map[keyType]valueType{} creates a new map with the specified key and value types.
		- map[keyType]valueType{key1: value1, key2: value2} creates a new map with the specified key and value types and initializes it with the given key-value pairs.
		- map[keyType]valueType{key1: value1, key2: value2, ...} creates a new map with the specified key and value types and initializes it with the given key-value pairs.
		- make(map[keyType]valueType) creates a new map with the specified key and value types.
		- map[keyType]valueType{key1: value1, key2: value2, ...} creates a new map with the specified key and value types and initializes it with the given key-value pairs.
		- make(map[keyType]valueType) creates a new map with the specified key and value types.
	*/

	person := map[string]string{
		"name":    "Naufal",
		"address": "Surabaya",
	}

	fmt.Println(person)            // Printing the entire map
	fmt.Println(person["name"])    // Accessing value by key
	fmt.Println(person["address"]) // Accessing value by key

	person["name"] = "Dzaki"    // Updating value for an existing key
	fmt.Println(person["name"]) // Now it prints "Dzaki"

	/*
		Alternative ways to create a map
		make(map[keyType]TypeValue) creates a new map with the specified key and value types.
	*/
	book := make(map[string]string)  // Creating an empty map
	book["title"] = "Belajar Golang" // Adding a key-value pair
	book["author"] = "Naufal Dzaki"  // Adding another key-value pair
	fmt.Println(book)                // Printing the map

	/*
		delete(map, key) removes the key-value pair from the map.
		If the key does not exist, it does nothing.
	*/
	book["whop"] = "asdasdasd"
	fmt.Println(book)    // Before deletion
	delete(book, "whop") // Deleting the key "whop"
	fmt.Println(book)    // After deletion, "whop" key is removed
}
