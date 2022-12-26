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
	aditya.NewEmail() // this method won't actually change the email.. because it passes a copy , and that copy is changed
	fmt.Println(aditya)
	fmt.Printf("the details are : %+v \n", aditya)
	fmt.Printf("name is %v and the email is %v. \n", aditya.Name, aditya.Email)
	aditya.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("online : ", u.Status)
}

func (u User) NewEmail() {
	u.Email = "haha@haha.com"
}
