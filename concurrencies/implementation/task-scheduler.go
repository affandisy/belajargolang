package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func TaskSchedulerA(i int, wg *sync.WaitGroup, ch1 chan bool) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		time.Sleep(time.Duration(rand.Intn(100) * int(time.Millisecond)))
		fmt.Printf("Task A: %d \n", i)
	}

	ch1 <- true
}

func TaskSchedulerB(i int, wg *sync.WaitGroup, ch1, ch2 chan bool) {
	defer wg.Done()
	<-ch1
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Duration(rand.Intn(100) * int(time.Millisecond)))
		fmt.Printf("Task B: %d \n", i)
	}

	ch2 <- true
}

func TaskSchedulerC(i int, wg *sync.WaitGroup, ch2 chan bool) {
	defer wg.Done()
	<-ch2
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Duration(rand.Intn(100) * int(time.Millisecond)))
		fmt.Printf("Task C: %d \n", i)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup

	ch1 := make(chan bool)
	ch2 := make(chan bool)

	fmt.Println("Task Scheduler: ")

	wg.Add(1)
	go TaskSchedulerA(5, &wg, ch1)

	wg.Add(1)
	go TaskSchedulerB(5, &wg, ch1, ch2)

	wg.Add(1)
	go TaskSchedulerC(5, &wg, ch2)

	wg.Wait()
	time.Sleep(3 * time.Second)

	fmt.Println("Task Scheduler Completed")

}
