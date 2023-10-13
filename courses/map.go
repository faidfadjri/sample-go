package main

import "fmt"

// Map adalah tipe data object pada go lang ( key value )
// String pertama menentukan tipe data untuk key
// String kedua menentukan tipe data untuk value

func tryMap() {
	// # Cara Pertama
	var person = map[string]string{
		"firstName" : "Faid",
		"lastName" : "Fadjri",
	}

	person["address"] = "Banjarnegara, Jawa Tengah"

	// Cara akses datanya
	fmt.Println(person)
	fmt.Println(person["firstName"])
	fmt.Println(person["lastName"])


	// # Cara Kedua
	book := make(map[string]string) // Gunakan := saat awal ( inisiasi variable )
	book["title"] = "Belajar Map Golang"
	book["author"] = "Faid Fadjri"

	fmt.Println(book)
}