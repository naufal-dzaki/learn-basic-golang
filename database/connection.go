package database

/*
What is Package Initialization in Go?
- Package initialization refers to the automatic execution of init() functions when a package is imported.
- Go will run:
	- All variable declarations with initial values.
	- All init() functions in the order they appear in the source file.
- The init() function:
	- Is called automatically before main() runs.
	- Cannot be called manually.
	- Does not accept parameters and does not return anything.
- It’s useful for:
	- Setting up default values.
	- Initializing connections or resources.
	- Validating configuration before actual code execution.
*/

var connection string

func init() {
	connection = "MySQL"
} // init() sets its value to "MySQL" when the database package is imported.

func GetDatabase() string {
	return connection
}
