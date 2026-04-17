package main

import "fmt"

func main(){
	fmt.Print("float number 32,64")
	var a float32 = 3.14 // 
	var b float64 = 3.8183999393
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("Complex Numbers")
	var c complex64 = 1 + 2i
	var d complex128 = 1 + 2i
	fmt.Println(c)
	fmt.Println(d)
} 