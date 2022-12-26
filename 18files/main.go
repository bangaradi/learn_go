package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "this is going to go in file"
	file, err := os.Create("./writeFile.txt")

	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println("the length is: ", length)
	defer file.Close()
	readFile("./writeFile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("text data inside the file :", string(databyte))
}
