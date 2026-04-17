package main

import "fmt"

func main()  {
	fmt.Println("Struct")
	type Person struct{
		Name string
		Age int
	}
	var q Person = Person{"Alice", 30}
	fmt.Println("q",q)
	fmt.Println("Name=", q.Name)
	fmt.Println("Age=", q.Age)
}