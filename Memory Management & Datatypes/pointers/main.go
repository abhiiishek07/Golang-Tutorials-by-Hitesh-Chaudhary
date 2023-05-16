package main

import "fmt"

func main() {
	fmt.Println("welcome to pointers")
	//pointer is the direct refernce to the memeory address

	// var ptr *int;
	// fmt.Println("value of ptr is", ptr)  // nilfp

	num := 10
	var ptr = &num
	fmt.Println("value of ptr is ", ptr)  // we get address of memory address of num ie 0xc64987855
	fmt.Println("value of ptr is ", *ptr) // 10

	*ptr = *ptr * 2
	fmt.Println("val of num is", num) //20

	var ptr2 = &num
	fmt.Println("val of ptr2 is", *ptr2)

}
