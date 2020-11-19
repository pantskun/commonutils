package main

import (
	"fmt"
	"os"
)

func main() {
	panicFunc()
}

func panicFunc() {
	defer func() {
		fmt.Println("start defer func")

		if r := recover(); r != nil {
			fmt.Println("get panic", r)
		}
	}()

	fmt.Println("ready to panic")
	// panic("panic")
	os.Exit(0)
}
