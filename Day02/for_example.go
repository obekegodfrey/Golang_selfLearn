package main

import "fmt"

func main() {
	//fmt.Println("Basic for loop!")
	//for i := 0; i < 5; i++ {
	//	fmt.Println(i)
	//}
	//fmt.Println("For as a while loop")
	//i := 0
	//for i < 5 {
	//	fmt.Println(i)
	//	i++
	//}
	//fmt.Println("Infinite for loop")
	//j := 0
	//for {
	//	if j >= 20 {
	//		break // break out of the loop when j is 100 or more
	//	}
	//	fmt.Println(j)
	//	j++
	//}
	fmt.Println("For loop with range(for slices, maps,strings)")
	nums := []int{10, 20, 30, 40, 50}
	for i, v := range nums {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
}
