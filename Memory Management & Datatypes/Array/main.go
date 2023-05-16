package main

import "fmt"

func main() {
	fmt.Println("Welc0me to array")

	var fruitsList [4]string

	fruitsList[0] = "Apple"
	fruitsList[1] = "Mango"
	// fruitsList[2]="peach"
	fruitsList[3] = "grapes"

	fmt.Println("fruits List is : ", fruitsList) // if a index is missed then there will be a space apple mango   grapes
	fmt.Println("Length is", len(fruitsList))    //4

	var vegList = [3]string{"potato", "tomato", "beans"}
	fmt.Println("veg list is:", vegList)
}
