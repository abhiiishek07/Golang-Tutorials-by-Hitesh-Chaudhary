package main

import "fmt"

// jwtToken := 50545   It will throw error. Use var jwtToken int = 50545

func main() {
	var username string = "Abhishek"
	fmt.Println(username)
	fmt.Printf("Type of username is %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("type of isLoggedIn is %T \n", isLoggedIn)

	var smallVal int = 256
	fmt.Println(smallVal)
	fmt.Printf("type is %T \n", smallVal)

	var floatVal float32 = 256.0255456688
	fmt.Println(floatVal)
	fmt.Printf("type is %T \n", floatVal)

	// default value and aliases

	var anotherVar int
	fmt.Println(anotherVar) // 0
	fmt.Printf("type is %T \n", anotherVar)

	var anotherStr string
	fmt.Println(anotherStr) //      ""
	fmt.Printf("type is %T \n", anotherStr)

	// implicit type
	var name = "Abhishek"
	fmt.Println(name) // lexer will decide the type
	fmt.Printf("type is %T \n", name)

	// no var style ( walrus style)

	totalusers := 3000 // this style only allowed inside medtod/functions
	fmt.Println(totalusers)
	fmt.Printf("type is %T", totalusers)
}
