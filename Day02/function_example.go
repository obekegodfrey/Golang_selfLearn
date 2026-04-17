package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func swap(x string, y string) (string, string) {
	return y, x
}

func split(sum int) (x int, y int) {
	x = sum * 4 / 2
	y = sum - x
	return
}

func sums_of_num(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println("Basic Functions:")
	result := add(3, 4)
	fmt.Println("3 + 4 =", result)
}
