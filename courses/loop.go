package main

import "fmt"



func main(){
	counter := 1	

	for counter <= 10 {
		fmt.Println("Perulangan ke - ", counter)
		counter++
	}


	// For Statement
	// Init : statement sebelum perulangan
	// Post : statement setelah perulanagan

	// counter := 1 adalah init statement
	// counter++ adalah post statement
	fmt.Println("------------ FOR STATEMENT -------------------")
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Loop statement : ", counter)
	}


	// For range = perulangan sejumlah array (e.g foreach di PHP)

	fmt.Println("------------ FOR RANGE -------------------")
	slice := []string{"Mohamad", "Faid", "Fadjri"}
	
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}