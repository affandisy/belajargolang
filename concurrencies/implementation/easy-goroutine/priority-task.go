package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	Name     string
	Priority int
}

type Driver struct {
	Id   int
	Name string
}

func (d Driver) assignTask(taskChan <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range taskChan {
		fmt.Printf("Driver %s (%d): Processing Task: %s \n", d.Name, d.Id, task.Name)
		time.Sleep(3 * time.Second)
	}
}

func PriorityDispatcher(tasks []Task, taskChan chan<- Task) {
	for priority := 1; priority <= 3; priority++ {
		for _, task := range tasks {
			if task.Priority == priority {
				taskChan <- task
			}
		}
	}
	close(taskChan)
}

func main() {
	task := []Task{
		{"Emergency Medical Transport", 1},
		{"Delivery at Zone A", 2},
		{"Pickup at Zone B", 3},
	}

	taskChan := make(chan Task, len(task))

	var wg sync.WaitGroup

	PriorityDispatcher(task, taskChan)

	drivers := []Driver{
		{1, "Syihabuddin Affandi"},
		{2, "Momonga King"},
	}

	for _, driver := range drivers {
		wg.Add(1)
		go driver.assignTask(taskChan, &wg)
	}

	wg.Wait()
	fmt.Println("All task is completed")
}
