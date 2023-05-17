package main

import "fmt"

func main() {
	fmt.Println("welocme to Methods ")

	user := User{"Abhishek", "abhi@xyz", true, 21}
	fmt.Println(user)                             // {Abhishek abhi@xyz true 21}
	fmt.Printf("user details are : %+v\n", user)  // user details are : {Name:Abhishek Email:abhi@xyz Status:true Age:21}
	fmt.Printf("Name of user is %v\n", user.Name) // Abhishek
	user.GetStatus()
	user.NewMail()
	fmt.Printf("user details are : %v\n", user) // email is not changed... structs/method passes a copy
}

type User struct { // capital bcoz its gonna be publically used
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() { // method
	fmt.Println("Is user active ? : ", u.Status)
}
func (u User) NewMail() {
	u.Email = "newmail@gmail.com"
	fmt.Println("Email of user is : ", u.Email)
}
