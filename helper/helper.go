package helper

/*
What is a Package in Go?
- A package is a way to group related Go files together.
- Each Go file starts with a package declaration that defines which package it belongs to.
- Packages help organize code and promote reuse and modularity.
- The main package is special — it defines the entry point of the program (the main() function).
*/

func SayHello(name string) string {
	return "Hello " + name
}
