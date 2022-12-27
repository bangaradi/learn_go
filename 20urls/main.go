package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "htttps://adi.tya:3000/learn?coursename=learn_go&paymentid=324kj23h4d"

func main() {
	fmt.Println(myUrl)

	// parsing the url --> using the net/url package :tick:
	result, _ := url.Parse(myUrl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	fmt.Println("")

	qParams := result.Query()
	fmt.Println(qParams)

	fmt.Println("")

	// making a url from given data :tick:
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "adi.tya",
		Path:    "learn",
		RawPath: "user=adityaB",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
