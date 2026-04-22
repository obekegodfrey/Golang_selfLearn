package main

import (
	"fmt"
)

func main(){
	//Basic declaration of a struct
	type Person struct {
		Name string
		Age int
	}

	//Declaring and instantiating together
	john := Person(Name: "John", Age: 30)
	fmt.Println("Declared and Instantited together:", john)

    //Using the new keyword
	jane := new(Person)
	jane.Name = "Jane"
	jane.Age = 25
	fmt.Println("Using new Keyword:", *jane)

	//Anonymous structs
	anon := struct {
		Country string 
		code int
	}{
		Country: "USA",
		Code: 1,
	}
	fmt.Println("Anonymous struct:", anon)

	//Nested structs
	type Address struct {
		City string
		State string
	}
	type Employee struct {
		PersonalInfo Person
		Address Address
	}
	emp := Employee{
		PersonalInfo: Person(Name: "Alice", Age: 28),
		Address: Address(City: "New York", State: "NY")
	}
	fmt.Println("Nested structs:", emp)

	//Embedded(Anonymous) Fields
	type Manager struct {
		person
		Department string
	}

	mgr := Manager{
		person: Person(Name: "Bob", Age: 35)
		department: "Finance"
	}
	fmt.Println("Embedded fields:", mgr)
}