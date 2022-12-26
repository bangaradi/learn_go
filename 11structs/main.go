package main

import "fmt"

func main() {
	// there is no inheritance in GOLANG; also there is no PARENT or SUPER concept
	aditya := User{
		"Aditya",
		"some@mail.com",
		true,
		20,
	}

	fmt.Println(aditya)
	fmt.Printf("the details are : %+v \n", aditya)
	fmt.Printf("name is %v and the email is %v. \n", aditya.Name, aditya.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
