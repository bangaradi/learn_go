package main

import "fmt"

func main()  {
	// map[<dataype of key>]<datatype of value>
	myMap := make(map[string]string)
	myMap["JS"] = "JavaScript"
	myMap["RB"] = "Ruby"
	myMap["PY"] = "Python"

	fmt.Println("the map as its printed is:", myMap)

	// deleting from the map
	delete(myMap, "RB")
	fmt.Println("the map is:", myMap)

	// looping over the map
	for key, value := range myMap {
		fmt.Printf("key:%s --> value:%s \n", key, value)
	}

}