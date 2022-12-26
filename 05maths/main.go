package main

import (
	"fmt"
	"math/big"
	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("welcome to maths in go")

	// var myNumberOne int = 2
	// var myNumberTwo float64 = 4.5

	// fmt.Println("the sum is:", myNumberOne+int(myNumberTwo))

	// random number ..using the math/rand package
	// rand.Seed(time.Now().UnixNano()) //setting the seed
	// fmt.Println(rand.Intn(5))        // getting a random number between 0-5

	// random numbers using the cryto package
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println("rand number: ", myRandomNum)
}
