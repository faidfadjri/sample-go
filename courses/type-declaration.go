package main

import "fmt"

func main() {
	// Type declaration adalah kita buat "Alias" atau nama lain sebagai tipe data
	type isMarried bool

	var marriedStatus isMarried = false

	fmt.Println(marriedStatus)
}