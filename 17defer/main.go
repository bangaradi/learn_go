package main

import "fmt"

func main() {
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	fmt.Println("world")
	myDefer()
}

func myDefer() {
	defer fmt.Println("")
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

// all that defer does is that it pushes the execution of the method at the end of the function block
// if there are multiple methods called for differ, then they are executed at the end in the LIFO manner :tick:

// THE OUTPUT will be:

// world
// 43210
// three
// two
// one
