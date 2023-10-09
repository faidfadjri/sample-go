package main

import "fmt"

func main() {

	// String
	var firstName string = "Mohamad"
	var lastName string = "Fadjri"
	var middleName string

	middleName = "Faid"
	middleName = "Changes middleName"

	fmt.Println(firstName + middleName + lastName)

	// Simple way for initialize variable instead of the var otherName string = "Mohamad Faid Fadjri"
	// Only the first time
	otherName := "Mohamad Faid Fadjri"
	fmt.Println(otherName)

	// Int
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// Bits & memory
	var numOne int8 = 120
	fmt.Print(numOne)
}
