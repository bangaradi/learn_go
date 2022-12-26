package main

import "fmt"

func main() {
	days := []string{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i:= range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v \n", index, day)
	// }

	randomVal := 10

	for randomVal > 5 {

		if randomVal == 9 {
			goto labelname
		}

		if randomVal == 7 {
			// break
			randomVal--
			continue
		}

		fmt.Println("running the loop value is ", randomVal)
		randomVal--
	}

labelname:
	fmt.Println("we hit the label name !!")
	fmt.Println("this is line 2")
}
