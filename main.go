package main

import (
	"fmt"
	"time"
)

func saySomething(name string) {
	fmt.Println("Hi", name)
}

func main() {
	// Start a Goroutine to run the saySomething function concurrently
	go saySomething("This is go routine")

	// Call saySomething without Goroutine, so it runs in the main thread
	saySomething("Fellas")

	// Sleep for a short time to allow Goroutines to complete
	time.Sleep(100 * time.Millisecond)

	fmt.Println("Main function has finished executing.")
}
