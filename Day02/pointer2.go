package day02
package main

import "fmt"

//Employee struct represent an employee record
type Employee struct{
	Name string
	Position String
	Salary float64
}

//UpdateSalary updates the salary of an employee
func UpdateSalary(emp *Employee, newSalary float64){
	emp.Salary = newSalary
}

func main()  {
	// Create an employee instance
	emp := Employee{
		Name: "John Doe",
		Position: "Software Engineer",
		Salary: 60000,
	}

	fmt.Printf("Before salary update: %v\n", emp)

	// Update the employee's salary
	UpdateSalary(&emp, 75000)

	fmt.Printf("After salary update: %v\n", emp) 
}