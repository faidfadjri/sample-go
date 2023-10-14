package main

import "fmt"

// Panic
// Program yang berfungsi untuk mengentikan program jika terjadi error
// Namun saat panic di eksekusi defer akan tetap berjalan

// Recover
// Function untuk menangkap data panic
// Jadi aplikasi masih tetap berjalan

func main(){
	runApp(true)
}	


func endApp(){
	fmt.Println("End Application...")


	message := recover()
	if message != nil {
		fmt.Println("Error Issue :",message)
	}
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Error")
	}
}