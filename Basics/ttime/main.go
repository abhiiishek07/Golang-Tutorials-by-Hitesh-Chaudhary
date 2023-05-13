package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time in go")

	curTime := time.Now()
	fmt.Println(curTime)

	fmt.Println(curTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2023, time.May, 13, 20, 25, 32, 12, time.UTC)
	fmt.Println(createdDate)
}
