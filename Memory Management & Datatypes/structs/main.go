package main

import "fmt"

func main() {
	fmt.Println("welocme to structs  ")

	user := User{"Abhishek", "abhi@xyz", true, 21}
	fmt.Println(user)                             // {Abhishek abhi@xyz true 21}
	fmt.Printf("user details are : %+v\n", user)  // user details are : {Name:Abhishek Email:abhi@xyz Status:true Age:21}
	fmt.Printf("Name of user is %v\n", user.Name) // Abhishek
}

type User struct { // capital bcoz its gonna be publically used
	Name   string
	Email  string
	Status bool
	Age    int
}
