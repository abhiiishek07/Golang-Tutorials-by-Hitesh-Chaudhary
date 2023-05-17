package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to functions in go")
	fmt.Println("Enter 1st num")
	reader := bufio.NewReader(os.Stdin)
	input1, _ := reader.ReadString('\n')
	// num1, _ := strconv.ParseInt(strings.TrimSpace(input1), 10, 0)
	num1, _ := strconv.Atoi(strings.TrimSpace(input1))
	fmt.Println("num 1 is", num1)

	fmt.Println("Enter 2nd num")
	input2, _ := reader.ReadString('\n')
	// num2, _ := strconv.ParseInt(strings.TrimSpace(input2), 10, 0)
	num2, _ := strconv.Atoi(strings.TrimSpace(input2))
	fmt.Println("Sum is", add(num1, num2))

	res2, msg := addAllNum(5, 6, 9, 7, 2, 5, 7, -2, 23, 5, 5, 6, 45, 3, 25)
	fmt.Println(msg, res2)
}
func add(num1 int, num2 int) int {
	return num1 + num2
}
func addAllNum(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Sum of all nums"
}
