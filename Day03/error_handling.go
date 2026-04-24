package main

import (
	"errors"
	"fmt"
)

// CustomError is a custom error type
type customError struct {
	msg string
}

// Error implements the error interface for custmError
func (e *customError) Error() string {
	return e.msg
}

// Somefunction demonstrates returning a custom error
func SomeFunction(flag boot) error {
	if !flag {
		//Return a custom error
		return &customError{"Custom error occurred"}
	}
	return nil
}

// SafeFunction demonstrates using panic and recover
func SafeFunction() {
	defer func() {
		if r := recover(); r != nil {

		}
	}()

	//This would cause a panic
	panic("A problem occurred")
}

// divide performs division and handles division by zero error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	//Simple error return
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Another example with valid division
	result, err = divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("result", result)
	}

	// Example of checking and handling errors with custom errors:
	err2 := SomeFunction(false)
	if err != nil {
		fmt.Println("Error:", err2)
	}

	//Example of using panic and recover
	SafeFunction()

}
