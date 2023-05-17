package main

import "fmt"

func main() {
	defer fmt.Println("world")  // defer will put this line in the end
	defer fmt.Println("first")  // LIFO
	defer fmt.Println("Second") // output will be : Hello Second First World
	fmt.Println("Hello")
	myDefer()
}
func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

/*
  Hello
  4
  3
  2
  1
  0
  Second
  First
  world
*/
