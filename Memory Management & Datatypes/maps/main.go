package main

import "fmt"

func main() {
	fmt.Println("welcome to maps")

	freq := make(map[string]int)

	freq["A"] = 5
	freq["B"] = 3
	freq["c"] = 15

	fmt.Println("freq of characters", freq)
	fmt.Println("freq of B is", freq["B"])

	delete(freq, "B")
	fmt.Println("new freq is", freq)

	//looping through maps

	for key, value := range freq {
		fmt.Printf("Freq of char %v is %v\n", key, value)
	}

}
