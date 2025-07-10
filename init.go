package main

/*
What is a Blank Identifier Import (_) in Go?
- import _ "some/package" tells Go to import and initialize the package without using any of its exported names.
- This triggers the package’s init() function even if you don’t reference anything from it.
- Commonly used when:
	- You need side effects from the init() function (e.g., logging, registration, initialization).
	- You want to execute setup code without directly accessing exported functions or variables.
*/

// When main imports the database package, Go runs the init() function inside database.
import (
	// "belajar-golang-dasar/internal" // error because imported and not used
	"belajar-golang-dasar/database" // The _ tells Go: “Import this package just for its side effects.”
	_ "belajar-golang-dasar/internal"
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}
