package main

import "fmt"

// Closure (Lingkup kerja sebuah function atau variable)
// Kemampuan sebuah function berinteraksi dengan data data di sekitarnya dalam scope yang sama


func Closure(){
	counter := 0
	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()

	fmt.Println(counter)
}