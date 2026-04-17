package main

import "fmt"
 
func main()  {
	fmt.Println("Signed Integers")
	var a int8 = 127 // ranges from -128 to 127
	var b int16 = -32768 // ranges from -32768 to 32767
	fmt.Println("a", a)
	fmt.Println("b", b)

	fmt.Println("Unsigned Integers")
	var c unint8 = 255 // ranges from 0 to 255
	var d unint16 = 65535 // ranges from 0 to 65535
	fmt.Println("c", c)
	fmt.Println("d", d)

	fmt.Println("Machine Dependent Types")
	var e  int = 123456789
	var f unit = 123456789
	var g unitptr = 0Xdeadbeef
	fmt.Println("e", e)
	fmt.Println("f", f)
	fmt.Println("g", g)

}