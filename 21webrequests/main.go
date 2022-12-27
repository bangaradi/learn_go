package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// GetRequest()
	// PostJsonRequest()
	PostFormRequest()
}

func GetRequest() {
	const myurl = "http://localhost:1111/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println(response)
	fmt.Println(response.StatusCode)
	fmt.Println(response.ContentLength)

	var responseString strings.Builder

	content, _ := ioutil.ReadAll(response.Body)
	// fmt.Println(string(content))

	byteCount, _ := responseString.Write(content)
	fmt.Println(byteCount)
	fmt.Println(responseString.String())
}

func PostJsonRequest() {
	const myurl = "http://localhost:1111/post"

	//fake json payload
	requestBdoy := strings.NewReader(`
		{
		"coursename": "learn_g0",
		"price": 0,
		"platform" : "youtube"	
		}
	`)

	// using http post in GO -> params are (url, content-type, request body)
	response, err := http.Post(myurl, "application/json", requestBdoy)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(response)
	fmt.Println(string(content))

}

func PostFormRequest() {
	const myurl = "http://localhost:1111/post"

	// making form data
	data := url.Values{}
	data.Add("firstname", "aditya")
	data.Add("lastname", "bangar")
	data.Add("email", "aditya@laksfdj.asdklfk")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println(response)
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
