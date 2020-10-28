package main

import "fmt"

func main() {
	fmt.Println("Please input a number")

	var i int
	fmt.Scanf("%d\n", &i)

	fmt.Println("input number: ", i)

	fmt.Println("Please input a string")

	var f float32
	fmt.Scanf("%f\n", &f)

	fmt.Println("input float: ", f)
}
