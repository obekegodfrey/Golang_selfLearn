package main
package main

import "fmt"

func main() {
	fmt.Println("Channel example in Go:")
	u := make(chan int) // unbuffered channel of integers
	fmt.Println("u:", u)
}