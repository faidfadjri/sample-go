package main

// Slice sama seperti array tapi size nya bisa berubah
// Slice juga array sebenarnya

import "fmt"

func main() {
	// ... uknown capacity
	// Membuat slice dari array
	var months = [...]string {
		"January",
		"February",
		"March", 
		"April",
		"May", 
		"June", 
		"July", 
		"August", 
		"September", 
		"October", 
		"November", 
		"December",
	}


	// Mengambil data dari index ke 4 dan index sebelum 7
	var slice1 = months[4:7]
	fmt.Println(slice1);


	// Mengambil data setelah index 10
	var slice2 = months[10:]
	fmt.Println(slice2)

	// Ketika di append dan slice sudah penuh maka akan membuat array baru
	var slice3 = append(slice2, "Faid");
	fmt.Println(slice3)

	slice3[1] = "Bukan December";
	fmt.Println(slice2)
	fmt.Println(slice3)

	// Membuat slice dari awal

	// 2 adalah length slice,, 5 capacity / value
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Faid"
	newSlice[1] = "Fadjri"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
}