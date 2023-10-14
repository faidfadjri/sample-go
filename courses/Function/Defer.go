package main

import "fmt"

// Defer / Trigerred Function
// Sebuah function yang dijadwalkan akan dirunning setelah suatu function selesai dijalankan

func Defer() {
	RunApp()
}

func RunApp() {
	defer logging()
	fmt.Println("Run Application...")
}

func logging(){
	fmt.Println("Setelah function selesai di eksekusi")
}