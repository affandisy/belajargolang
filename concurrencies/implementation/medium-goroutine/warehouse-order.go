package main

import (
	"fmt"
	"sync"
	"time"
)

type OrderWarehouse struct {
	Name     string
	Priority int
}

type WorkerWarehouse struct {
	Name string
	Id   int
}

func (worker WorkerWarehouse) assignOrder(orderChan <-chan OrderWarehouse, wg *sync.WaitGroup) {
	defer wg.Done()

	for order := range orderChan {
		time.Sleep(time.Duration(1+order.Priority) * time.Second)

		fmt.Printf("Worker %s (%d): Processed Order: %s [Priority %d] \n", worker.Name, worker.Id, order.Name, order.Priority)
	}
}

func OrderDispatcher(orders []OrderWarehouse, orderChan chan<- OrderWarehouse) {
	for priority := 1; priority <= 3; priority++ {
		for _, order := range orders {
			if order.Priority == priority {
				orderChan <- order
			}
		}
	}
	close(orderChan)
}

func main() {
	orderWarehouses := []OrderWarehouse{
		{"Laptop", 1},   // Express
		{"Mouse", 2},    // Regular
		{"Keyboard", 2}, // Regular
		{"Monitor", 3},  // Bulk
		{"Printer", 1},  // Express
	}

	workerWarehouses := []WorkerWarehouse{
		{"Syihabuddin Affandi", 1},
		{"Gibs Gajah", 2},
		{"Jhon Koo Weeh", 3},
		{"Happy Seventeen", 4},
		{"Istimewa", 5},
	}

	var wg sync.WaitGroup

	orderChan := make(chan OrderWarehouse)
	// resultChan := make(chan string)

	for _, worker := range workerWarehouses {
		wg.Add(1)
		go worker.assignOrder(orderChan, &wg)
	}

	go OrderDispatcher(orderWarehouses, orderChan)

	wg.Wait()

	fmt.Println("All Orders processed Successfully")
}
