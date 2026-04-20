package main

import (
	"fmt"
)

func main() {
	//Initialize a map
	fruits := map[string]int{"apple": 5, "banana": 7}

	//Check length
	fmt.Println("Length of the map:", len(fruits))

	//Adding Elements
	fruits["orange"] = 10
	fmt.Println("Added 'orange':", fruits)

	//Retieving Elements
	applePrice := fruits["apple"]
	fmt.Println("Price of apple:", applePrice)

	// Checking Existence
	price, exists := fruits["kiwi"]
	if exists {
		fmt.Println("Price of kiwi:", price)
	} else {
		fmt.Println("Kiwi doesnot exist in the map")
	}

	orangePrice, exists := fruits["orange"]
	if exists {
		fmt.Println("price of orange", orangePrice)
	} else {
		fmt.Println("Orange doesnot exist in the map")
	}

	// Deleting Elements
	delete(fruits, "banana")
	fmt.Println("After deleting 'banana':", fruits)

}
