package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Welcome to maths in go")

	//using math/rand
	// r := rand.New(rand.NewSource(time.Now().Unix()))
	// fmt.Println(r.Intn(132))

	//usinf crypto/rand

	randNum, _ := rand.Int(rand.Reader, big.NewInt(7))
	fmt.Println(randNum)

}
