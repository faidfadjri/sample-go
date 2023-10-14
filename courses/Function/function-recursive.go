package main

import "fmt"

// Recursive function
// Function yang memanggil dirinya sendiri
// Factorial 5! = 5 * 3 * 2 * 1

func RecursiveFunction(){
	fmt.Println(Factorial(4))
}

func Factorial(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * Factorial(value - 1)
	}

}