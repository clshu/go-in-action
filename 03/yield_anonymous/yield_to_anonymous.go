package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Outside goroutine")
	go func() {
		fmt.Println("Inside a goroutine")
	}()
	fmt.Println("Outside again")

	runtime.Gosched()
}
