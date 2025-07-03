package main

import "fmt"

func main() {
	/*
		NoKTP is a custom type based on string
		you can create a custom type to represent a specific kind of data
		like NoKTP (Nomor KTP) which is a string but has a specific meaning
	*/
	type NoKTP string

	var KTPNaufal NoKTP = "1234567890123456"
	var kTPString string = "123456789012111"
	var KTP NoKTP = NoKTP(kTPString) // you can convert string to NoKTP type
	fmt.Println(KTPNaufal)
	fmt.Println(KTP)
}
