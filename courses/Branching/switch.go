package main

import "fmt"


func main(){
	name := "Faid"

	switch name {
		case "Faid":
			fmt.Println("Nama Depan")
		case "Fadjri":
			fmt.Println("Nama Belakang")
		default:
			fmt.Println("Tidak diketahui")		
	}


	// Short statement

	fmt.Println("Short Statement")
	switch length := len(name); length > 3 {
		case true:
			fmt.Println("Karakter lebih dari 3")
		default:
			fmt.Println("Karakter kurang dari 3")
	}


	// Switch tanpa expression
	age := 10
	switch {
		case age > 5 :
			fmt.Println("Anak TK")
		default :
			fmt.Println("Balita")
	}
}