package helper

/*
What is a Package in Go?
- A package is a way to group related Go files together.
- Each Go file starts with a package declaration that defines which package it belongs to.
- Packages help organize code and promote reuse and modularity.
- The main package is special — it defines the entry point of the program (the main() function).
*/

/*
What is an Access Modifier in Go?
- In Go, access control (also called access modifier) is determined by the first letter of an identifier.
- Go uses capitalization to control access level — not special keywords like in other languages.
- Capitalized Identifiers → Exported (Public/Can be accessed from other packages.)
- Lowercase Identifiers → Unexported (Private/Only accessible within the same package.)
*/

var version = "1.0.0"      // unexported — lowercase, cannot be accessed from outside
var Application = "golang" // exported — capitalized, can be accessed from other packages

func sayGoodBye(name string) string { // unexported
	return "Good bye " + name
}

func SayHello(name string) string { // exported
	return "Hello " + name
}
