package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("give me some input :")

	// the input err syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("the input was, ", input)
	fmt.Printf("type of the input was %T", input)

}
