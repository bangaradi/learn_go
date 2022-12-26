package main

import "fmt"

func main() {
	// var ptr *int
	// fmt.Println("the value of the pointer is:", ptr)
	//gives null as output

	Number := 21
	var ptr *int = &Number

	fmt.Println("value of the pointer: ", ptr)
	fmt.Println("value the pointer is pointing towards: ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("new value of pointer :", *ptr)
}
