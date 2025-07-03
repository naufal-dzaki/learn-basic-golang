package main

import "fmt"

func main() {
	name1 := "naufal"
	name2 := "naufal"
	isSameName := name1 == name2
	isNotSameName := name1 != name2

	fmt.Println(name1)
	fmt.Println(name2)
	fmt.Println(isSameName)
	fmt.Println(isNotSameName)
}
