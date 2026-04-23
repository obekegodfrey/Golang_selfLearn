package main

import (
	"fmt"
)

// 1. Single-method Interface
type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + " says Woof!"
}

// 2. Embedding Interface
type Animal interface {
	Speaker
	Move() string
}

type Bird struct {
	Name string
}

func (b Bird) Speak() string {
	return b.Name + " says Tweet!"
}

func (b Bird) Move() string {
	return b.Name + " files"
}

// 3. Empty Interface
func PrintAnything(v interface{}) {
	fmt.Println(v)
}

// 4. Interface with Multiple Methods
type Vehicle interface {
	Start() string
	Stop() string
}

type car struct {
	Model string
}

func (c Car) start() string {
	return c.Model + " car stopped"
}

//5. Functional interface(single Method Interface)
//Similar to Speaker Interface

//6. Interface as a Constraint(Generics)
func Describe[T Speaker](t T) {
	fmt.Println(t.Speak())
}

func main(){
	dog := Dog(Name: "Buddy")
	bird := Bird(Name: "Tweety")
	car := Car(Model: "Tesla")

	fmt.println(dog.speak())
	fmt.Println(bird.speak())
	fmt.Println(bird.Move())

	PrintAnything("An empty interface can hold anything")

	fmt.println(car.Star())
	fmt.Println(car.stop())

	Describe(dog)
	Describe(bird)
}

