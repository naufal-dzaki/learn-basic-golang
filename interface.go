package main

import "fmt"

/*
interface
- An interface is a type that defines a contract or a set of methods that a struct must implement.
- It allows you to define behavior without specifying how that behavior is implemented.
- A type implements an interface implicitly by defining all the methods declared in the interface — no need to explicitly declare implementation.
- Interfaces allow you to write flexible, reusable, and decoupled code.
*/

type HashName interface {
	GetName() string
}

func SayHello(name HashName) {
	fmt.Println("Hello", name.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{"Naufal Dzaki"}
	SayHello(person)

	animal := Animal{"cat"}
	SayHello(animal)
}
