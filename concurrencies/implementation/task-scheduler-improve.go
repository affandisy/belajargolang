package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func RunTask(label string, wg *sync.WaitGroup, start <-chan bool, done chan<- bool) {
	defer wg.Done()

	if start != nil {
		<-start
	}

	for i := 1; i <= 5; i++ {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Printf("Task %s: %d \n", label, i)
	}

	if done != nil {
		done <- true
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup

	startB := make(chan bool)
	startC := make(chan bool)

	fmt.Println("Task Scheduler: ")

	wg.Add(1)
	go RunTask("A", &wg, nil, startB)

	wg.Add(1)
	go RunTask("B", &wg, startB, startC)

	wg.Add(1)
	go RunTask("C", &wg, startC, nil)

	wg.Wait()
	fmt.Println("Task Scheduler Completed")
}
