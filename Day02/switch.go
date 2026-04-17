package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Switch statement example")
	//day := 3
	//switch day {
	//case 1:
	//	fmt.Println("Monday")
	//case 2:
	//	fmt.Println("Tuesday")
	//case 3:
	//	fmt.Println("Wednesday")
	//case 4:
	//	fmt.Println("Thursday")
	//case 5:
	//	fmt.Println("Friday")
	//case 6:
	//	fmt.Println("Saturday")
	//case 7:
	//	fmt.Println("Sunday")
	//default:
	//	fmt.Println("Invalid day")
	//}
	//fmt.Println("Switch statement with initialization")
	//switch today := time.Now().Weekday(); today {
	//case time.Saturday, time.Sunday:
	//	fmt.Println("Today is the weekend")
	//default:
	//	fmt.Println("Today is a weekday")
	//}
	fmt.Println("Switch without a condition{like if-else}")
	x := 42
	switch {
	case x < 0:
		fmt.Println("x is negative")
	case x == 0:
		fmt.Println("x is zero")
	case x > 0 && x < 50:
		fmt.Println("x is a positive number less than 50")
	default:
		fmt.Println("x is a positive number greater than or equal to 50")
	}
}
