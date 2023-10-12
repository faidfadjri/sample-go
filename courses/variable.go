package main

import "fmt"



/*

----------------------------------
var int8Var int8 = -128 // Ini adalah deklarasi tipe data int8 dengan nilai minimal -128 hingga maksimal 127.
var int16Var int16 = -32768 // Ini adalah deklarasi tipe data int16 dengan nilai minimal -32768 hingga maksimal 32767.
var int32Var int32 = -2147483648 // Ini adalah deklarasi tipe data int32 dengan nilai minimal -2147483648 hingga maksimal 2147483647.
var int64Var int64 = -9223372036854775808 // Ini adalah deklarasi tipe data int64 dengan nilai minimal -9223372036854775808 hingga maksimal 9223372036854775807.


var uintVar uint = 127 // Ini adalah deklarasi tipe data uint dengan nilai minimal 0 hingga maksimal 255 (tergantung pada arsitektur platform).
var float32Var float32 = 3.14 // Ini adalah deklarasi tipe data float32.
var float64Var float64 = 3.141592653589793 // Ini adalah deklarasi tipe data float64.
*/

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
