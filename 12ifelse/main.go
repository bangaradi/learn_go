package main

import "fmt"

func main() {

	loginCount := 15

	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 && loginCount < 20 {
		result = "middle"
	} else {
		result = "high"
	}

	fmt.Println(result)

}
