package main

import "fmt"

func main() {
	fmt.Println("Welocme to looping in GO ")

	// days := []string{"sun", "mon", "tue", "wed", "thr", "fri", "sat"}

	// for day := 0; day < len(days); day++ {
	// 	fmt.Println(days[day])
	// }

	// for day := range days {
	// 	fmt.Println(days[day])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index of %v is %v\n", day, index)
	// }

	i := 1
	for i < 10 { // while loop
		if i == 3 {
			goto lco
		}
		if i == 7 {
			fmt.Println("hbd abhishek")
			break
		}
		fmt.Println(i)
		i++
	}
lco:
	fmt.Println("Hello")

}
