package main

import (
	"fmt"
)

var C int = 0;
var done int = 0;

func P1() {
	// Increment C 10k times
	for i := 0; i < 10000; i++ {
		C = C + 1
	}
	// Signal that this thread is done
	done++
}

func main() {
	for i := 0; i < 100; i++ {
		C = 0
		done = 0
		// Run 2 threads
		go P1()
		go P1()
		// Wait until both are done 
		for done != 2 {}
		// Print result
		fmt.Println("C = ", C)
	}
}