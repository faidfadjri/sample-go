package main

import "fmt"

// Membuat Alias (Type Declaration)
// Harus sebuah function yang return nya sebagai Boolean
type Blacklist func(string) bool


func main() {
	// Anonymous Function
	// Membuat function yang di save ke variable
	fmt.Println("----- ANONYMOUS FUNCTION ------")


	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("Faid Fadjri", blacklist)
}


func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}