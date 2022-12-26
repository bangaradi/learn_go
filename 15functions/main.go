package main

import "fmt"

func main() {

	//simple function call
	//greeter()

	// passing function as a reference
	// refToGreeter := greeter
	// refToGreeter()

	//function to add 2 numbers
	result := adder(3, 5)
	fmt.Println("the result is :", result)

	// making a function with arbitray number of inputs
	total := arbAdder(2, 3, 44, 5, 5, 6, 7, 7, 8)
	fmt.Println("result of arbAdder:", total)

}

func adder(one int, two int) int {
	return one + two
}

func arbAdder(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}
func greeter() {
	fmt.Println("hello this is a greeting")
}
