package main

import (
	"fmt"
)

// Implementing an interface
type Greeter interface {
	Greet() string
}

type Person struct {
	Name string
}

func (p Person) Greet() string {
	return "Hello, " + p.Name
}

// Type Assertions
func printDetails(i interface{}) {
	str, ok := i.(string)
	if ok {
		fmt.Println("it's a string:", str)
	} else {
		fmt.Println("Not a string")
	}
}

// Interface with type Switches
func Describe(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("Unknown type")
	}

}

func main() {
	//Implementing an interface
	person := Person{Name: "Alice"}
	var greeter Greeter = person
	fmt.Println(greeter.Greet())

	// Type assertions
	printDetails("Hello")
	printDetails(42)

	//Interface with Type Switches
	Describe("hello, world")
	Describe(2023)
	Describe(3.14)
}
