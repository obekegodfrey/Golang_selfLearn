package main

import (
	"fmt"
	"regexp"
)

func main() {
	//Sample inputs to available
	name := "Obeke Godfrey"
	email := "obekegodfrey2@gmail.com"
	phoneNumber := "+256-701-263-898"

	// Regular expression patterns
	namePattern := `^[a-zA-Z]+(?: [a-zA-Z]+)*$`
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	phonePattern := `^\+\d{1,3}-\d{3}-\d{3}-\d{3}$`

	//compile regrex and match patterns
	matchname, _ := regexp.MatchString(namePattern, name)
	matchEmail, _ := regexp.MatchString(emailPattern, email)
	matchPhone, _ := regexp.MatchString(phonePattern, phoneNumber)

	// Print results
	if matchname {
		println("Name is valid.")
	} else {
		println("Name is invalid.")
	}

	if matchEmail {
		println("Email is valid.")
	} else {
		println("Email is invalid.")
	}

	if matchPhone {
		println("Phone number is valid.")
	} else {
		println("Phone number is invalid.")
	}

	// output results
	fmt.Printf("Name: %s, Valid: %t\n", name, matchname)
	fmt.Printf("Email: %s, Valid: %t\n", email, matchEmail)
	fmt.Printf("Phone Number: %s, Valid: %t\n", phoneNumber, matchPhone)

}
