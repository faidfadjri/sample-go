// Channel adalah tempat komunikasi synchronous dalam go routine
// Channel memiliki pengirim & penerima (TX & RX)

// Mekanisme Channel
// Ketika kita membuat go routine dan mendefinisikan channel
// Go routine akan mengirim data ke channel.
// Channel akan menunggu sampai ada "go routine penerima" mengambil data dari channel yang dikirim
// e.g async -> await di javascript

// Secara default channel hanya bisa menampung satu data. jadi satu channel satu data
// Tetapi data dari channel bisa diambil lebih dari 1 go routine.
// Channel harus di close jika tidak digunakan agar tidak terjadi memory leak.

package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {

	channel := make(chan string)
	defer close(channel) // jika function TestCreateChannel sudah selesai maka akan jalan code ini

	go func() {
		time.Sleep(time.Second)
		// Mengirim data ke channel
		channel <- "Faid Fadjri"
		fmt.Println("Selesai mengirim data ke channel")
		fmt.Println("Waiting...")

		// Jika data tidak diambil maka channel akan stuck disini
	}()

	// Jika data diambil tapi ternyata tidak ada data maka akan error "deadlock!"
	data := <-channel

	fmt.Println("Mengambil data dari channel...")
	fmt.Println(data)
}
