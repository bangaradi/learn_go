package main

import "fmt"

func main() {
	var fruits [4]string
	fruits[0] = "apple"
	fruits[1] = "pineapple"
	// fruits[2] = "kiwi"
	fruits[3] = "guava"

	fmt.Println("the fruits are :", fruits)
	fmt.Println("the length of fruits: ", len(fruits))

	var vegetables= [3]string{"potato", "beans", "eggplant"}
	fmt.Println("the list of vegetables is :", vegetables)

}
