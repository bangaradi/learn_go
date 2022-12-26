package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"strconv"
)

func main() {
	fmt.Println("welcome to our pizza app")
	fmt.Println("thanks for eating our pizza, please rate: (b/w 1- 5)")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating: ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added 1 to your rating: ", numRating+1)
	}

}
