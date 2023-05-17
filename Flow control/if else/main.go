package main

import "fmt"

func main() {
	fmt.Println("welcome to if else")

	loginCnt := 23
	var result string

	if loginCnt < 10 {
		result = "regular user"
	} else if loginCnt >= 10 && loginCnt < 50 {
		result = "not so regular"
	} else {
		result = "watch out"
	}
	fmt.Println(result)

	if num := 11; num%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}
