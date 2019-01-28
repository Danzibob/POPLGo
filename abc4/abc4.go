package main

import (
	"fmt"
	"os"
	"time"
)

var f_next [3]bool

func P(letter string, position int){
	for {
		
		for !f_next[position] {}
		
		f_next[position] = false

		for i := 0; i < 4; i++{
			fmt.Print(letter)
		}

		nextPos := (position + 1) % 3
		f_next[nextPos] = true
	}
}

func main() {
	f_next = [3]bool{true,false,false}
	
	go P("A",0)
	go P("B",1)
	go P("C",2)

	// Terminate the program after 40ms
	time.Sleep(40*time.Millisecond)
	fmt.Println("\nExiting")
	os.Exit(0)
}