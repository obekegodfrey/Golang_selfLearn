package main

import (
	"fmt"
	"reflect"
)

// Person struct definitions
type Person struct {
	Name string
	Age int
}

//method on the Person struct
func (p *Person) SayHello() {
	fmt.Println("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

//Employee struct wwith tagged fields
type Employee struct {
	ID int 'json:"id"'
	Name string 'json:"name"'
}

func main(){
  // Create an instance of a struct
  alice := Person{Name: "Alice", Age: 30}
  fmt.println("Struct Instance:", alice)

   //Accessing and modifying fileds
   fmt.Println("Name before:", alice.Name)
   alice.Name = "ALice Smith"
   fmt.Println("NAme after:", alice.Name)
  
   //Pointers to structs
   bob := &Person{Name: "Bob", Age:25}
   fmt.Println("Pointer to Struct:", bob)
   
   //Methods on structs
   alice.SayHello()
   bob.SayHello() 

   //Tagging struct fields
   emp := Employee{ID: 1, Name: "John Doe"}
   t := reflect.TypeOf(emp)
   field, _ := t.FieldByName("Name")
   fmt.Println("Tag on Name filed:", field.Tag.Get("json "))
   }




