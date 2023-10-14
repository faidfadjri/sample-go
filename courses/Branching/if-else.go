package main

import "fmt"



func tryIfElseCondition() {
	name := "Faid"

	if name == "Faid" {
		fmt.Println("Benar")
	} else {
		fmt.Println("Salah")
	}


	// Short statement
	// Deklarasi sesuatu sebelum kondisi -> deklarasikan variable dulu menentukan kondisi if-else nya

	if length := len(name); length > 5 {
		fmt.Println("Karakter lebih dari 5")
	} else {
		fmt.Println("Karakter kurang dari 5")
	}

}