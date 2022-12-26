package main

import "fmt"

// the walrus operator doesn't work outside a function or a method
// someVar := 1312  -> this is wrong is here

// IMP note
const LoginToken = "haokdfjs" // note that the first letter of this variable is CAPITAL
// this is the syntax used to declare public variables... this variable can be accessed by other files under the same package

func main() {
	var username string = "aditya"
	fmt.Println(username)
	fmt.Printf("variable is of the type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of the type: %T \n", isLoggedIn)

	//default values and aliases
	var unitializedVar int
	fmt.Println(unitializedVar)

	//implicit type declaration
	var implicitVar = "heello"
	fmt.Println(implicitVar)

	//shorthand or no var style
	noVarStyle := 34.4
	fmt.Println(noVarStyle)

}
