package main

import (
	"fmt"
	"time"
)

// Representation of forks
// True = fork on table
// Phil. 0 is sat between forks 0 (left) and 1 (right)
// Phil. 1 is sat between forks 1 (left) and 2 (right)
// Phil. 2 is sat between forks 2 (left) and 0 (right)
var forks [3]bool
var eating = false


func Philosopher(name string, pos int){
	rpos := (pos+1)%3
	// Representation of Phil.'s hands
	left,right := false,false
	// Try to pick up both forks
	// This is where out concurrency issue arises
	left,right,forks[pos],forks[rpos] = forks[pos],forks[rpos],left,right
	
	for {
		// Check if we're holding both forks
		if left && right {
			// If we are, eat
			fmt.Println(name,"eats...")
			time.Sleep(500 * time.Millisecond)
			// Then put forks down
			forks[pos], forks[rpos], left, right = true, true, false, false
			// Rest so we don't pick up the forks again
			time.Sleep(1 * time.Millisecond)
		} else if right {
			// Wait until left fork is available
			for !forks[pos] {}
			forks[pos], left = left, forks[pos]
		} else {
			// Wait until right fork is available
			for !forks[rpos] {}
			forks[rpos], right = right, forks[rpos]
		}
	}
}

func main() {
	forks = [3]bool{true,true,true}
	
	go Philosopher("Fred",0)
	go Philosopher("Bob",1)
	go Philosopher("Jim",2)
	for {}
}