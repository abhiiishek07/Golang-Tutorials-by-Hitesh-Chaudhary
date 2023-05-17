package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switch case")

	rand.NewSource(time.Now().Unix())
	// diceNum := rand.Intn(6) + 1
	diceNum := 4
	fmt.Println(diceNum)

	switch diceNum {
	case 1:
		fmt.Println("Dice val is 1")
	case 2:
		fmt.Println("Dice num is 2")
	case 3:
		fmt.Println("Dice num is 3")
	case 4:
		fmt.Println("Dice num is 4")
		fallthrough // when dice val is 4 it will print case 5 also
	case 5:
		fmt.Println("Dice num is 5")

	case 6:
		fmt.Println("Dice num is 6")
	default:
		fmt.Println("Invalid")
	}
}
