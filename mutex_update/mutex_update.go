package main

import (
	"fmt"
	"time"
)

var flag1, flag2 bool
var C = 0;

func P1() {
	for {
		
		flag1 = true
		for flag2 {}
		C = C + 1
		flag1 = false

	}
}

func P2() {
	for {
		
		flag2 = true
		for flag1 {}
		C = C + 1
		flag2 = false

	}
}


func main() {
	go P1()
	go P2()
	for {
		time.Sleep(500*time.Millisecond)
		fmt.Println(C)
	}
}