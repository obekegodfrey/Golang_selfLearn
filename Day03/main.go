package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <number>")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n < 0 {
		fmt.Println("Please provide a non-negative integer.")
		return
	}

	fmt.Printf("Factorial of %d is %d\n", n, factorial(n))
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}
