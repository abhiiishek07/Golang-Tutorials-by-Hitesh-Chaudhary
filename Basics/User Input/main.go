package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Hello welcom to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the pizza rating: ")

	// comma ok or error ok syntax... its like try & catch

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating", input)

}
