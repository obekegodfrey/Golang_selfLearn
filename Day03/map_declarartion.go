package main

import (
	"fmt"
)

func main() {
	// using the var keyword without initializing it
	var map1 map[string]int
	fmt.Println("Map declared with var keyword without initialization:", map1)

	// Using the map Literal
	map2 := map[string]int{"apple": 5, "banana": 7}
	fmt.Println("Map declared with a map literal:", map2)

	// using the make function
	map3 := make(map[string]int)
	map3["apple"] = 5
	map3["banana"] = 7
	fmt.Println("Map declared with make function:", map3)

	// using the make function with specific size
	map4 := make(map[string]int, 10) // 10 is the initial size
	map4["apple"] = 5
	map4["banana"] = 7
	fmt.Println("map declared with make function and specific size:", map4)

	// Using Structs as map values
	type Fruit struct {
		Price int
		Color string
	}
	map5 := make(map[string]Fruit)
	map5["apple"] = Fruit{Price: 5, Color: "green"}
	map5["banana"] = Fruit{Price: 7, Color: "orange"}
	fmt.Println("Map with structs as values:", map5)

	//Nested Maps

	map6 := make(map[string]map[string]int)
	map6["fruit"] = map[string]int{"apple": 5, "banana": 7}
	map6["fruit"] = map[string]int{"apple": 5, "banana": 7}
	fmt.Println("Nested map:", map6)
}
