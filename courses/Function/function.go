package main

import "fmt"


func FirstFunction(){
	first, middle, last := getCompleteName("Mohamad", "Faid", "Fadjri")
	fmt.Println(first, " ", middle, " ", last)


	fmt.Println("----- VARIADIC FUNCTION ------")
	fmt.Println(sumAll(5, 20, 40, 50))


	numbers := []int{10,20,30,40}
	fmt.Println(sumAll(numbers...))




	// Function Value
	// Membuat semacam object / instance bukan dari class tapi dari function

	fmt.Println("----- FUNCTION VALUE ------")

	sumAllFunction := sumAll
	fmt.Println(sumAllFunction(10,30,40))

	// Function as parameter
	sayHello("Anjing",spamFilter)


	
}

func getCompleteName(firstName, middleName, lastName string) (string, string, string) {
	firstName = "Mohamad"
	middleName = "Faid"
	lastName = "Fadjri"
	return firstName, middleName, lastName
}


// Variadic function
// Parameter terakhir bisa berfungsi sebagai varargs
// varargs data bisa menerima satu input (semacam array)

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// Function as parameter
// passing function sebagai parameter (e.g : callback function)

func sayHello(name string, filter func(string) string) {
	fmt.Println("Hello ", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}