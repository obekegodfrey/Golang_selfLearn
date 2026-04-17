package day02
package main

import "fmt"

func main() {
    fmt.Println("function as variable")

	var v func(int) int
	v = func(x int) int {
		retrun x * x
	}
	result := v(5) // result will be 25
	fmt.Println("result:", result)
}