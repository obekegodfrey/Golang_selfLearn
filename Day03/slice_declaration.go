package main

import "fmt"

func main() {
	//1. using the var keyword without initializing
	var slice1 []int
	fmt.Println("Slice1 {unintialized}:", slice1)

	//Initialize slice1 before using it
	slice1 = []int{1, 2, 3}
	fmt.Println("Slice1 {initialized}:", slice1)

	//2. Using a slice literal
	slice2 := []int{4, 5, 6}
	fmt.Println("Slice2:", slice2)

	//3. using the make function
	slice3 := make([]int, 3) //Length and capacity are 3
	fmt.Println("Slice 3 before putting variable", slice3)
	slice3[0] = 7
	slice3[1] = 8
	slice3[2] = 9
	fmt.Println("Slice3:", slice3)
	slice4 := make([]int, 3, 5)
	slice4[0] = 7
	slice4[1] = 8
	slice4[2] = 9
	fmt.Println("Slice 4:", slice4)
	slice4 = append(slice4, 10)
	slice4 = append(slice4, 11)
	fmt.Println("Slice 4 after append:", slice4)
	fmt.Println("Slice 4 capacity:", cap(slice4))
	slice4 = append(slice4, 12)
	slice4 = append(slice4, 13)
	fmt.Println("Slice 4 capacity:", cap(slice4))
	fmt.Println("Slice 4 after append:", slice4)

	//4. From an Array
	array := [5]int{10, 11, 12, 13, 14}
	slice5 := array[2:4]
	fmt.Println("Slice4:", slice5)

}
