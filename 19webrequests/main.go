package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("response is of type: %T \n", response)

	defer response.Body.Close() // it's the caller's responsibility to close the connection

	databyte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databyte)
	fmt.Print(content)
}
