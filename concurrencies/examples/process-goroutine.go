package main

import (
	"fmt"
	"runtime"
	"time"
)

func FirstProcess(index int) {
	fmt.Println("First process func Started")
	for i := 1; i <= index; i++ {
		fmt.Println("i= ", i)
	}
	fmt.Println("First process func Ended")
}

func SecondProcess(index int) {
	fmt.Println("Second process func started")
	for j := 1; j <= index; j++ {
		fmt.Println("j= ", j)
	}
	fmt.Println("Second process func Ended")
}

func main() {
	fmt.Println("Main Execution Started")

	go FirstProcess(8)

	SecondProcess(8)

	fmt.Println("No. of Goroutines: ", runtime.NumGoroutine())

	time.Sleep(time.Second * 3)

	fmt.Println("Main Execution Ended")
}
