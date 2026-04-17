package main

import "fmt"

func main() {
	var matrix [3][3]int

	//Initialize the metrix
	count := 1
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = count
			count++
		}
	}

	// Print the matrix
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println("%d", matrix[i][j])
		}
		fmt.Println()
	}
}
