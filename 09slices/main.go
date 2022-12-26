package main

import (
	"fmt"
	"sort"
)

// the SLICES data structure is more popular and powerful in go as compared to arrays
// it has many methods and it has dynamic memory allocation

func main() {
	var fruits = []string{"apple", "kiwi", "peach"}
	fmt.Printf("the datatype of fruits is: %T \n", fruits)

	// the append method
	fruits = append(fruits, "mango", "grapes")
	fmt.Println("the fruits are :", fruits)

	// slicing up our slice
	fruits = append(fruits[1:])
	fmt.Println("the fruits list are slicing is:", fruits)

	// using the MAKE keyword to allocate memory
	numbers := make([]int, 4)
	numbers[0] = 23
	numbers[1] = 53
	numbers[2] = 33
	numbers[3] = 12

	//numbers[4] = 99 --> ERROR
	// but the append method works
	numbers = append(numbers, 32, 11, 22 ,33)
	fmt.Println("the numbers are:", numbers)

	// sorting the numbers
	// we make use of the sort package
	sort.Ints(numbers)
	fmt.Println("the sorted array of numbers is: ", numbers)

	// how to remove a value from a slices based on index
	var courses = []string{"reactjs", "js", "ruby", "python", "swift"}
	var index int = 2
	
	// we want to delete the value at index number 2
	courses = append(courses[:index], courses[index+1:]...)
	// the ... is to tell the function to unpack the slice and append it to the given slice, i.e., courses[:index]
}
