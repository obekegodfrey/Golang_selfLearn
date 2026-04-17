package main

import "fmt"

func main() {
	fmt.Println("function as variable")

	var v func(int) int
	v = func(x int) int {
		return x * x // This function takes an integer x and returns its square
	}
	result := v(5) // result will be 25
	fmt.Println("result:", result)
}
