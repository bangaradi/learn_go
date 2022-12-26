package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	dice := rand.Intn(6) + 1
	fmt.Println("the value of the dice is: ", dice)

	switch dice {
	case 1:
		fmt.Println("dice value is 1 and you can open")
	case 2:
		fmt.Println("its 2 ")
	case 3:
		fmt.Println("its 3 ")
	case 4:
		fmt.Println("its 4 ")
		fallthrough // fallthrough is used when we want to run the next case as well.....
		// in similar fashion we can have multiple fall throughs so that after one case multiple other cases can run...
	case 5:
		fmt.Println("its 5 ")
	case 6:
		fmt.Println("its 6 and you can roll the dice again ")
	default:
		fmt.Println("what the hell was that..")
	}


}
