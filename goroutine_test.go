// Runing test
// go test -v -run=TestCreateGoRoutine

// Go Routine -> Thread ringan sebuah task yang berjalan didalam thread.
// Go Routine tidak berjalan secara bersamaan (parallel) tapi bergantian tergantung task yang paling cepat
// Secara default ketika menggunakan go routine. go akan skip task yang berpotensi lama. ( maka task tersebut tidak dijalankan )
// Secara default memanggil function dengan go routine tidak dapat menggunakan function yang memiliki return value

package main

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello world!")
}

func TestCreateGoRoutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoRoutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
