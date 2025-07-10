package main

/*
What is Import in Go?
- The import keyword is used to include code from other packages.
- This allows you to use functions, types, and variables defined elsewhere.
- You can import:
	- Standard library packages like fmt, math, etc.
	- Your own custom packages, like your helper package.
*/

import (
	"belajar-golang-dasar/helper" // project name/package name
	"fmt"
)

func main() {
	result := helper.SayHello("Naufal")
	fmt.Println(result)
}
