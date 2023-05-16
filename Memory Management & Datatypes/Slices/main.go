package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slices")

	var fruitsList = []string{"Apple", "tomato", "guava"}
	// fmt.Printf("type is %T\n", fruitsList) // []string

	fruitsList = append(fruitsList, "Mango", "grapes") //adding fruits to list
	// fmt.Println(fruitsList)

	// fruitsList = append(fruitsList[2:]) // start from index till last index
	// fmt.Println(fruitsList)             // guava, mango, grapes

	// fruitsList = append(fruitsList[1:4])  // start from index 1 till 3
	// fmt.Println(fruitsList)            //tomato guava grapes

	fruitsList = append(fruitsList[:3]) //start from index 0 till 2
	// fmt.Println(fruitsList)             // Apple tomato guava

	scores := make([]int, 5)
	scores[0] = 5
	scores[1] = 3
	scores[2] = 65
	scores[3] = 0
	scores[4] = 99
	// scores[5] = 100         index out of bound error
	// fmt.Println(scores)

	scores = append(scores, 100, 102, 69) // it will allocate memory
	// fmt.Println(scores)                   // 5 3 65 0 99 100 102 69

	sort.Ints(scores)
	// fmt.Println(scores)
	// fmt.Println(sort.IntsAreSorted(scores))

	// Remove a value from slices based on index

	var mySkills = []string{"React", "Go", "Node", "Mongo DB", "firebase", "MUI", "DSA"}
	var index int = 2 // index to be removed
	mySkills = append(mySkills[:index], mySkills[index+1:]...)
	fmt.Println(mySkills)

}
